package db

import (
	"context"
	//TODO: perhaps migrate belongs better in storage
	_ "github.com/lib/pq"
	"github.com/sudipidus/pismo-test/logger"
	"github.com/sudipidus/pismo-test/models"
	storage2 "github.com/sudipidus/pismo-test/storage"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func SeedOperationType(storage2 storage2.Storage) {
	opTypes := []models.OperationType{
		{
			ID:          1,
			Type:        "CASH_PURCHUASE",
			Description: "cash Purchuase using account",
			IsCredit:    false,
		},
		{
			ID:          2,
			Type:        "INSTALLMENT_PURCHUASE",
			Description: "installment purchuase from account",
			IsCredit:    false,
		},
		{
			ID:          3,
			Type:        "WITHDRAWAL",
			Description: "withdrawal the account",
			IsCredit:    false,
		},
		{
			ID:          4,
			Type:        "PAYMENT",
			Description: "Payment to the account",
			IsCredit:    true,
		},
	}
	_, err := storage2.SeedOperationType(context.Background(), opTypes)
	if err != nil {
		logger.GetLogger().Fatal("Seeding of operation type data failed.. aborting", zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err})
	}
	_ = opTypes
	logger.GetLogger().Info("Seeding of operation type data complete")

}
