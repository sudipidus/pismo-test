package storage

import (
	"context"
	"database/sql"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sudipidus/pismo-test/errors"
	"github.com/sudipidus/pismo-test/models"
	"log"
	"math/rand"
	"time"
)

//go:generate mockgen -source=storage.go -destination=./mock_storage/mock_storage.go
type Storage interface {
	CreateAccount(context context.Context, account *models.Account) (*models.Account, *errors.Error)
	FetchAccount(ctx context.Context, accountID string) (*models.Account, *errors.Error)
	CreateTransaction(ctx context.Context, transaction *models.Transaction) (*models.Transaction, *errors.Error)
	UpdateTransactionBalance(ctx context.Context, transactionID int, balance float64) (*models.Transaction, *errors.Error)
	SeedOperationType(ctx context.Context, operationTypes []models.OperationType) (*[]models.OperationType, *errors.Error)
	FetchPendingTransaction(ctx context.Context, accountID int) ([]*models.Transaction, *errors.Error)
}

type PostgresStorage struct {
	Db *sqlx.DB
}

type InMemoryStorage struct {
	inmemoryStorage map[interface{}]interface{}
}

func NewPostgresStorage(dsn string) (*PostgresStorage, *errors.Error) {
	db, err := sqlx.Connect("postgres", dsn)

	//Connection pooling
	db.SetMaxOpenConns(1000) // The default is 0 (unlimited)
	db.SetMaxIdleConns(10)   // defaultMaxIdleConns = 2
	db.SetConnMaxLifetime(0) // 0, connections are reused forever.
	if err != nil {
		return nil, errors.NewError(500, "Db Setup Failed", err)
	}
	return &PostgresStorage{Db: db}, nil
}

func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{inmemoryStorage: make(map[interface{}]interface{})}
}

func (ps *PostgresStorage) CreateAccount(context context.Context, account *models.Account) (*models.Account, *errors.Error) {
	var newAccount models.Account
	//todo: negate amount for debit
	// todo: timestamps to have proper values
	query := `
        INSERT INTO accounts (
            document_number,
            created_at,
            updated_at
        ) VALUES (
            :document_number,
            :created_at,
            :updated_at
        )
        RETURNING id, document_number, created_at, updated_at
    `

	rows, err := ps.Db.NamedQuery(query, account)
	if err != nil {
		return nil, errors.NewError(500, "Failed to execute query", err)
	}
	defer rows.Close()

	if rows.Next() {
		if err := rows.StructScan(&newAccount); err != nil {
			return nil, errors.NewError(404, "Record not found", err)
		}
	}

	return &newAccount, nil
}

func (ps *PostgresStorage) FetchAccount(context context.Context, id string) (*models.Account, *errors.Error) {
	var account models.Account
	//todo: fetch by document number or (account nuber & Db id are separate, don't tie it)
	query := `
        SELECT id, document_number, created_at, updated_at
        FROM accounts
        WHERE id = :id
    `

	rows, err := ps.Db.NamedQuery(query, map[string]interface{}{"id": id})
	if err != nil {
		return nil, errors.NewError(500, "Failed to execute query", err)
	}
	defer rows.Close()

	if rows.Next() {
		if err := rows.StructScan(&account); err != nil {
			return nil, errors.NewError(500, "Error while preparing response", err)
		}
	} else {
		return nil, errors.NewError(404, "Account not found", err)
	}

	return &account, nil
}

func (ps *PostgresStorage) CreateTransaction(ctx context.Context, transaction *models.Transaction) (*models.Transaction, *errors.Error) {
	tx, ok := ctx.Value("tx").(*sql.Tx)
	if !ok {
		return nil, errors.NewError(500, "Transaction not found in context", nil)
	}

	var newTransaction models.Transaction
	query := `
        INSERT INTO transactions (
            account_id,
            operation_type_id,
            amount,
            balance,
            transaction_date,
            created_at,
            updated_at
        ) VALUES (
            $1,
            $2,
            $3,
            $4,
            $5,
        	$6,
        	$7          
        )
        RETURNING id, account_id, operation_type_id,amount,balance,transaction_date, created_at, updated_at
    `

	_, err := tx.ExecContext(ctx, query,
		transaction.AccountID,
		transaction.OperationTypeID,
		transaction.Amount,
		transaction.Balance,
		transaction.TransactionDate,
		transaction.CreatedAt,
		transaction.UpdatedAt)

	if err != nil {
		return nil, errors.NewError(500, "Failed to execute query", err)
	}

	// Commenting as last supported ID is not supported by this driver
	//id, err := result.LastInsertId()
	//if err != nil {
	//	return nil, errors.NewError(500, "Failed to get last insert id", err)
	//}

	//newTransaction.TransactionID = int(id)
	newTransaction.AccountID = transaction.AccountID
	newTransaction.OperationTypeID = transaction.OperationTypeID
	newTransaction.Amount = transaction.Amount
	newTransaction.Balance = transaction.Balance
	newTransaction.TransactionDate = transaction.TransactionDate
	newTransaction.CreatedAt = transaction.CreatedAt
	newTransaction.UpdatedAt = transaction.UpdatedAt

	return &newTransaction, nil
}

func (ps *PostgresStorage) UpdateTransactionBalance(ctx context.Context, txnID int, balance float64) (*models.Transaction, *errors.Error) {
	tx, ok := ctx.Value("tx").(*sql.Tx)
	if !ok {
		return nil, errors.NewError(500, "Transaction not found in context", nil)
	}

	query := `
        UPDATE transactions
        SET balance = $1
        WHERE id = $2
        RETURNING id, account_id, operation_type_id, balance, transaction_date, created_at, updated_at
    `

	row := tx.QueryRow(query, balance, txnID)
	if row == nil {
		return nil, errors.NewError(404, "Transaction not found", nil)
	}

	var txn models.Transaction
	err := row.Scan(&txn.TransactionID, &txn.AccountID, &txn.OperationTypeID, &txn.Balance, &txn.TransactionDate, &txn.CreatedAt, &txn.UpdatedAt)
	if err != nil {
		return nil, errors.NewError(500, "Failed to scan row", err)
	}

	return &txn, nil
}

func (ps *PostgresStorage) SeedOperationType(ctx context.Context, operationTypes []models.OperationType) (*[]models.OperationType, *errors.Error) {
	tx, err := ps.Db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	//defer func(tx *sql.Tx) {
	//	err := tx.Rollback()
	//	if err != nil {
	//		logger.GetLogger().Fatal("Failed in rolling back of transaction during seeding of operation data")
	//	}
	//}(tx)

	for _, opType := range operationTypes {
		_, err := tx.Exec(`
			INSERT INTO operation_types (id,type, description, is_credit, created_at, updated_at)
			VALUES ($1, $2, $3, $4, $5,$6)
			ON CONFLICT (id) DO NOTHING;
		`, opType.ID, opType.Type, opType.Description, opType.IsCredit, time.Now(), time.Now())
		if err != nil {
			log.Fatal(err)
		}
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
	return nil, nil
}

func (ps *PostgresStorage) FetchPendingTransaction(ctx context.Context, accountID int) ([]*models.Transaction, *errors.Error) {
	// transactions for accountID where type ID Is 1,2,3 sorted by created_at asc
	var pendingTransactions []*models.Transaction

	var transaction models.Transaction
	query := `
        SELECT id, account_id, operation_type_id, amount, balance, transaction_date, created_at, updated_at
        FROM transactions
        WHERE account_id = :account_id AND operation_type_id in (1,2,3) 
        and balance < 0
        order by created_at asc 
    `

	rows, err := ps.Db.NamedQuery(query, map[string]interface{}{"account_id": accountID})
	if err != nil {
		return nil, errors.NewError(500, "Failed to execute query", err)
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.StructScan(&transaction); err != nil {
			return nil, errors.NewError(500, "Error while preparing response", err)
		}
		pendingTransactions = append(pendingTransactions, &transaction)
	}

	return pendingTransactions, nil
}

func (ps *PostgresStorage) BeginTx(ctx context.Context) (*sql.Tx, error) {
	return ps.Db.Begin()
}

func (ms *InMemoryStorage) CreateAccount(context context.Context, account *models.Account) (*models.Account, *errors.Error) {
	account.ID = rand.Int()
	ms.inmemoryStorage[account.DocumentNumber] = account
	return account, nil
}

func (ms *InMemoryStorage) FetchAccount(context context.Context, accountID string) (*models.Account, *errors.Error) {
	return nil, nil
}

func (ms *InMemoryStorage) CreateTransaction(context context.Context, transaction *models.Transaction) (*models.Transaction, *errors.Error) {
	return nil, nil
}

func (ms *InMemoryStorage) SeedOperationType(ctx context.Context, operationTypes []models.OperationType) (*[]models.OperationType, *errors.Error) {
	return nil, nil
}

func (ms *InMemoryStorage) FetchPendingTransaction(ctx context.Context, accountID int) ([]*models.Transaction, *errors.Error) {
	return nil, nil
}

func (ms *InMemoryStorage) UpdateTransactionBalance(ctx context.Context, transactionID int, balance float64) (*models.Transaction, *errors.Error) {
	return nil, nil
}
