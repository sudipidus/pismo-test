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
			Type:        "NORMAL_PURCHUASE",
			Description: "Normal Purchuase using account",
			IsCredit:    false,
		},
		{
			ID:          2,
			Type:        "WITHDRAWAL",
			Description: "Withdrawal from account",
			IsCredit:    false,
		},
		{
			ID:          3,
			Type:        "CREDIT_VOUCHER",
			Description: "Voucher Credit to the account",
			IsCredit:    true,
		},
		{
			ID:          4,
			Type:        "INSTALLMENT_PURCHUASE",
			Description: "Installment Purchuase using account",
			IsCredit:    false,
		},
	}
	_, err := storage2.SeedOperationType(context.Background(), opTypes)
	if err != nil {
		logger.GetLogger().Fatal("Seeding of operation type data failed.. aborting", zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err})
	}
	_ = opTypes
	logger.GetLogger().Info("Seeding of operation type data complete")

}
