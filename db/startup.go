package db

import (
	"github.com/sudipidus/pismo-test/errors"
	storage2 "github.com/sudipidus/pismo-test/storage"
	"log"
	"os"
)

var storage storage2.Storage

func Init() {
	dsn := os.Getenv("DB_DSN")
	var err *errors.Error
	storage, err = storage2.NewPostgresStorage(dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v, aborting....", err)
	}
}

func GetStorage() storage2.Storage {
	return storage
}
