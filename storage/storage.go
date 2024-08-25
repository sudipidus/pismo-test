package storage

import (
	"context"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sudipidus/pismo-test/models"
	"github.com/sudipidus/pismo-test/serviceErrors"
	"math/rand"
)

//go:generate mockgen -source=storage.go -destination=./mock_storage/mock_storage.go
type Storage interface {
	CreateAccount(context context.Context, account *models.Account) (*models.Account, *serviceErrors.ServiceError)
	FetchAccount(ctx context.Context, accountID string) (*models.Account, *serviceErrors.ServiceError)
}

type PostgresStorage struct {
	db *sqlx.DB
}

type InMemoryStorage struct {
	inmemoryStorage map[interface{}]interface{}
}

func NewPostgresStorage(dsn string) (*PostgresStorage, *serviceErrors.ServiceError) {
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, &serviceErrors.ServiceError{WrappedError: err, Code: "INIT_ERROR"}
	}
	return &PostgresStorage{db: db}, nil
}

func NewInMemoryStorage() (*InMemoryStorage, *serviceErrors.ServiceError) {
	return &InMemoryStorage{inmemoryStorage: make(map[interface{}]interface{})}, nil
}

func (ps *PostgresStorage) CreateAccount(context context.Context, account *models.Account) (*models.Account, *serviceErrors.ServiceError) {
	var newAccount models.Account
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
		return nil, &serviceErrors.ServiceError{WrappedError: err, Code: "CREATE_ERROR"}
	}
	defer rows.Close()

	if rows.Next() {
		if err := rows.StructScan(&newAccount); err != nil {
			return nil, &serviceErrors.ServiceError{WrappedError: err, Code: "CREATE_ERROR"}
		}
	}

	return &newAccount, nil
}

func (ps *PostgresStorage) FetchAccount(context context.Context, id string) (*models.Account, *serviceErrors.ServiceError) {
	var account models.Account
	query := `
        SELECT id, document_number, created_at, updated_at
        FROM accounts
        WHERE id = :id
    `

	rows, err := ps.db.NamedQuery(query, map[string]interface{}{"id": id})
	if err != nil {
		return nil, &serviceErrors.ServiceError{WrappedError: err, Code: "FETCH_ERROR"}
	}
	defer rows.Close()

	if rows.Next() {
		if err := rows.StructScan(&account); err != nil {
			return nil, &serviceErrors.ServiceError{WrappedError: err, Code: "FETCH_ERROR"}
		}
	} else {
		return nil, &serviceErrors.ServiceError{WrappedError: err, Code: "NOT_FOUND"}
	}

	return &account, nil
}

func (ms *InMemoryStorage) CreateAccount(context context.Context, account *models.Account) (*models.Account, error) {
	account.ID = rand.Int()
	ms.inmemoryStorage[account.DocumentNumber] = account
	return account, nil
}
