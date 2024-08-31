package services

import (
	"context"
	"github.com/sudipidus/pismo-test/errors"
	"github.com/sudipidus/pismo-test/logger"
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
}

func NewPismoService(storage storage.Storage) PismoService {
	return &PismoServiceImpl{storage: storage}
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

	txn, err := s.storage.CreateTransaction(ctx, &models.Transaction{
		AccountID:       request.AccountID,
		Amount:          request.Amount,
		OperationTypeID: request.OperationTypeID,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	})
	if err != nil {
		return nil, err
	}
	return txn, nil
}
