package storage_test

import (
	"context"
	"fmt"
	storage2 "github.com/sudipidus/pismo-test/storage"
	"log"
	"os"
	"testing"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/sudipidus/pismo-test/models"
)

var db_dsn = "pismo-user-test:pismo-secret-test@127.0.0.1:5434/pismo-test?sslmode=disable"

func TestMain(m *testing.M) {
	// Set up a test database
	db, err := sqlx.Connect("postgres", "host=localhost port=5434 user=pismo-user-test password=pismo-secret-test dbname=pismo-test sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create the test database schema
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS accounts (
			id SERIAL PRIMARY KEY,
			document_number VARCHAR(255) NOT NULL,
			created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);

		CREATE TABLE IF NOT EXISTS transactions (
			id SERIAL PRIMARY KEY,
			account_id INTEGER NOT NULL,
			operation_type_id INTEGER NOT NULL,
			amount DECIMAL(10, 2) NOT NULL,
			transaction_date DATE NOT NULL,
			created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);

		CREATE TABLE IF NOT EXISTS operation_types (
			id SERIAL PRIMARY KEY,
			type VARCHAR(255) NOT NULL,
			description VARCHAR(255) NOT NULL,
			is_credit BOOLEAN NOT NULL,
			created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);
	`)
	if err != nil {
		log.Fatal(err)
	}

	// Run the tests
	code := m.Run()

	// Tear down the test database
	_, err = db.Exec("DROP TABLE IF EXISTS accounts, transactions, operation_types")
	if err != nil {
		log.Fatal(err)
	}

	os.Exit(code)
}

func TestPostgresStorage_CreateAccount(t *testing.T) {
	storage, err := storage2.NewPostgresStorage("postgres://" + db_dsn)
	assert.Nil(t, err)

	account := &models.Account{
		DocumentNumber: "123456789",
	}
	newAccount, err := storage.CreateAccount(context.Background(), account)
	assert.Nil(t, err)
	assert.NotNil(t, newAccount)
	assert.NotZero(t, newAccount.ID)
}

func TestPostgresStorage_FetchAccount(t *testing.T) {
	storage, err := storage2.NewPostgresStorage("postgres://" + db_dsn)
	assert.Nil(t, err)

	account := &models.Account{
		DocumentNumber: "123456789",
	}
	newAccount, err := storage.CreateAccount(context.Background(), account)
	assert.Nil(t, err)

	fetchedAccount, err := storage.FetchAccount(context.Background(), fmt.Sprintf("%d", newAccount.ID))
	assert.Nil(t, err)
	assert.NotNil(t, fetchedAccount)
	assert.Equal(t, newAccount.ID, fetchedAccount.ID)
}

func TestPostgresStorage_CreateTransaction(t *testing.T) {
	storage, err := storage2.NewPostgresStorage("postgres://" + db_dsn)
	assert.Nil(t, err)

	account := &models.Account{
		DocumentNumber: "123456789",
	}
	newAccount, _ := storage.CreateAccount(context.Background(), account)

	transaction := &models.Transaction{
		AccountID:       newAccount.ID,
		OperationTypeID: 1,
		Amount:          100.00,
		TransactionDate: time.Now(),
	}
	newTransaction, err := storage.CreateTransaction(context.Background(), transaction)
	assert.Nil(t, err)
	assert.NotNil(t, newTransaction)
	assert.NotZero(t, newTransaction.TransactionID)
}

func TestPostgresStorage_SeedOperationType(t *testing.T) {
	storage, err := storage2.NewPostgresStorage("postgres://" + db_dsn)
	assert.Nil(t, err)

	operationTypes := []models.OperationType{
		{
			Type:        "DEPOSIT",
			Description: "Deposit into account",
			IsCredit:    true,
		},
		{
			Type:        "WITHDRAWAL",
			Description: "Withdrawal from account",
			IsCredit:    false,
		},
	}
	_, err = storage.SeedOperationType(context.Background(), operationTypes)
	assert.Nil(t, err)
}
