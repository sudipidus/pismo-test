package services

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/sudipidus/pismo-test/errors"
	"github.com/sudipidus/pismo-test/logger"
	"math"
	"os"
	"time"

	"github.com/sudipidus/pismo-test/models"
	"github.com/sudipidus/pismo-test/storage"
)

//go:generate mockgen -source=PismoService.go -destination=./mock_services/mock_pismo_service.go
type PismoService interface {
	Greet(ctx context.Context) (interface{}, *errors.Error)
	CreateAccount(ctx context.Context, request CreateAccountRequest) (*models.Account, *errors.Error)
	FetchAccount(ctx context.Context, accountID string) (*models.Account, *errors.Error)
	CreateTransaction(ctx context.Context, request CreateTransactionRequest) (*models.Transaction, *errors.Error)
}

type PismoServiceImpl struct {
	storage storage.Storage
	lock    Lock
}

func NewPismoService(storage storage.Storage) PismoService {
	return &PismoServiceImpl{
		storage: storage,
		lock: NewRedisLock(redis.NewClient(&redis.Options{
			Addr: fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
		})),
	}
}

func (s *PismoServiceImpl) Greet(ctx context.Context) (interface{}, *errors.Error) {
	return "Greetings from Pismo-Test", nil
}

func (s *PismoServiceImpl) CreateAccount(ctx context.Context, request CreateAccountRequest) (*models.Account, *errors.Error) {

	account := &models.Account{
		DocumentNumber: request.DocumentNumber,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	account, err := s.storage.CreateAccount(ctx, account)
	if err != nil {
		logger.GetLogger().Error("account creation failed")
		return nil, err
	}
	return account, nil
}

func (s *PismoServiceImpl) FetchAccount(ctx context.Context, accountID string) (*models.Account, *errors.Error) {
	account, err := s.storage.FetchAccount(ctx, accountID)
	if err != nil {
		return nil, err
	}
	return account, nil
}

func (s *PismoServiceImpl) CreateTransaction(ctx context.Context, request CreateTransactionRequest) (*models.Transaction, *errors.Error) {
	accountID := request.AccountID
	lockKey := fmt.Sprintf("transaction_lock_for_account %d", accountID)
	err := s.lock.Lock(lockKey)
	if err != nil {
		fmt.Println("Error acquiring lock:", err)
		return nil, &errors.Error{Code: 500, Message: "Internal Server Error"}
	}
	defer s.lock.Unlock(lockKey)

	ps := s.storage.(*storage.PostgresStorage)
	tx, err := ps.BeginTx(ctx)
	defer tx.Rollback() // If it's already committed let it silently fail
	ctx = context.WithValue(ctx, "tx", tx)

	if err != nil {
		return nil, &errors.Error{Code: 500, Message: "Internal database error"}
	}

	txn, storageError := s.storage.CreateTransaction(ctx, &models.Transaction{
		AccountID:       request.AccountID,
		Amount:          multiplier(request.OperationTypeID) * request.Amount,
		Balance:         multiplier(request.OperationTypeID) * request.Amount,
		OperationTypeID: request.OperationTypeID,
		TransactionDate: time.Now(),
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	})
	if storageError != nil {
		return nil, &errors.Error{Code: 500, Message: "Internal database error"}
	}

	// Discharging process when there is credit (type 4 - payment)
	// discharge (pay) - for the accountID (oldest first, partial applicable)
	if request.OperationTypeID == 4 {
		msg := fmt.Sprintf("Discharging for account %d with %f amount for operation type %d", request.AccountID,
			request.Amount, request.OperationTypeID)
		logger.GetLogger().Info(msg)

		// fetch transactions belong to accountID sorted by created_at or debit type
		var pendingTransactions []*models.Transaction
		pendingTransactions, serviceErrr := getPendingTransactions(ctx, s.storage, request.AccountID)
		if serviceErrr != nil {
			logger.GetLogger().Error("failed to discharge")
			return nil, &errors.Error{Code: 500, Message: "Internal database error"}
		}

		paymentAmount := request.Amount
		for _, transaction := range pendingTransactions {
			if paymentAmount <= 0 {
				break
			}
			// balance is -50, paymentAmount is 100
			if paymentAmount >= math.Abs(transaction.Balance) {
				// since balance is already negative
				paymentAmount = paymentAmount + transaction.Balance
				transaction.Balance = 0
				saveTransaction(ctx, s.storage, transaction)
			} else {
				// balance is -50, paymentAmount is 25
				transaction.Balance = transaction.Balance + paymentAmount
				paymentAmount = 0
				saveTransaction(ctx, s.storage, transaction)
			}
		}
		txn.Balance = paymentAmount
		saveTransaction(ctx, s.storage, txn)
	}
	tx.Commit()
	return txn, nil
}

func saveTransaction(ctx context.Context, storage storage.Storage, transaction *models.Transaction) (*models.Transaction, *errors.Error) {
	return storage.UpdateTransactionBalance(ctx, transaction.TransactionID, transaction.Balance)
}

func getPendingTransactions(ctx context.Context, storage storage.Storage, accountID int) ([]*models.Transaction, *errors.Error) {
	return storage.FetchPendingTransaction(ctx, accountID)
}

func multiplier(operationTypeID int) float64 {
	if operationTypeID == 4 {
		return 1
	} else {
		return -1
	}
}
