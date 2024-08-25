package db

import (
	"context"
	_ "github.com/lib/pq"
	"github.com/sudipidus/pismo-test/logger"
	"github.com/sudipidus/pismo-test/models"
	storage2 "github.com/sudipidus/pismo-test/storage"
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
			Type:        "WITHDRAWAL",
			Description: "Withdrawal from account",
			IsCredit:    false,
		},
		{
			Type:        "CREDIT_VOUCHER",
			Description: "Voucher Credit to the account",
			IsCredit:    true,
		},
		{
			Type:        "INSTALLMENT_PURCHUASE",
			Description: "Installment Purchuase using account",
			IsCredit:    false,
		},
	}
	_, err := storage2.SeedOperationType(context.Background(), opTypes)
	if err != nil {
		logger.GetLogger().Fatal("Seeding of operation type data failed.. aborting")
	}
	logger.GetLogger().Info("Seeding of operation type data complete")

}
