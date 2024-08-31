package db

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/sudipidus/pismo-test/errors"
	"github.com/sudipidus/pismo-test/logger"
	storage2 "github.com/sudipidus/pismo-test/storage"
	"log"
	"os"
)

var storage storage2.Storage

const StorageTypeInMemory = "in-memory"
const StorageTypePostgres = "postgres"

func Init() {
	storageType := os.Getenv("STORAGE_TYPE")

	if storageType == StorageTypePostgres {
		dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
			os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
		// Run database migrations
		m, err := migrate.New(
			"file://db/migrations",
			dsn,
		)
		if err != nil {
			log.Fatalf("Failed to initialize migrations: %v", err)
		}

		err = m.Up()
		if err != nil && err != migrate.ErrNoChange {
			log.Fatalf("Failed to run migrations: %v", err)
		}

		// Initialize Postgres storage after migration
		var dbInitError *errors.Error
		storage, dbInitError = storage2.NewPostgresStorage(dsn)
		if dbInitError != nil {
			log.Fatalf("Failed to initialize Postgres storage: %v", dbInitError)
		}
	} else {
		storage = storage2.NewInMemoryStorage()
	}

	logger.GetLogger().Info("storage initialized with " + storageType)
}

func GetStorage() storage2.Storage {
	return storage
}
