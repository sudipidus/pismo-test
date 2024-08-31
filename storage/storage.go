package storage

import (
	"context"
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
	SeedOperationType(ctx context.Context, operationTypes []models.OperationType) (*[]models.OperationType, *errors.Error)
}

type PostgresStorage struct {
	db *sqlx.DB
}

type InMemoryStorage struct {
	inmemoryStorage map[interface{}]interface{}
}

func NewPostgresStorage(dsn string) (*PostgresStorage, *errors.Error) {
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, errors.NewError(500, "DB Setup Failed", err)
	}
	return &PostgresStorage{db: db}, nil
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

	rows, err := ps.db.NamedQuery(query, account)
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
	//todo: fetch by document number or (account nuber & db id are separate, don't tie it)
	query := `
        SELECT id, document_number, created_at, updated_at
        FROM accounts
        WHERE id = :id
    `

	rows, err := ps.db.NamedQuery(query, map[string]interface{}{"id": id})
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

func (ps *PostgresStorage) CreateTransaction(context context.Context, transaction *models.Transaction) (*models.Transaction, *errors.Error) {
	var newTransaction models.Transaction
	//todo: timestamps not set
	query := `
        INSERT INTO transactions (
            account_id,
            operation_type_id,
            amount,
            transaction_date,
            created_at,
            updated_at
        ) VALUES (
            :account_id,
            :operation_type_id,
            :amount,
            :transaction_date,
        	:created_at,
        	:updated_at          
        )
        RETURNING id, account_id, operation_type_id,amount,transaction_date, created_at, updated_at
    `

	rows, err := ps.db.NamedQuery(query, transaction)
	if err != nil {
		return nil, errors.NewError(500, "Failed to execute query", err)
	}
	defer rows.Close()

	if rows.Next() {
		if err := rows.StructScan(&newTransaction); err != nil {
			return nil, errors.NewError(404, "Record not found", err)
		}
	}

	return &newTransaction, nil
}

func (ps *PostgresStorage) SeedOperationType(ctx context.Context, operationTypes []models.OperationType) (*[]models.OperationType, *errors.Error) {
	tx, err := ps.db.Begin()
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
