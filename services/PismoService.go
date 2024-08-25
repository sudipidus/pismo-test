package services

import (
	"context"
	"github.com/sudipidus/pismo-test/logger"
	"github.com/sudipidus/pismo-test/serviceErrors"
	"time"

	"github.com/sudipidus/pismo-test/models"
	"github.com/sudipidus/pismo-test/storage"
)

//go:generate mockgen -source=PismoService.go -destination=./mock_services/mock_pismo_service.go
type PismoService interface {
	Greet(ctx context.Context) (interface{}, *serviceErrors.ServiceError)
	CreateAccount(ctx context.Context, request CreateAccountRequest) (interface{}, *serviceErrors.ServiceError)
	FetchAccount(ctx context.Context, accountID string) (interface{}, *serviceErrors.ServiceError)
	CreateTransaction(ctx context.Context) (interface{}, *serviceErrors.ServiceError)
}

type PismoServiceImpl struct {
	storage storage.Storage
}

func NewPismoService(storage storage.Storage) PismoService {
	return &PismoServiceImpl{storage: storage}
}

func (s *PismoServiceImpl) Greet(ctx context.Context) (interface{}, *serviceErrors.ServiceError) {
	return "Greetings from Pismo-Test", nil
}

func (s *PismoServiceImpl) CreateAccount(ctx context.Context, request CreateAccountRequest) (interface{}, *serviceErrors.ServiceError) {

	account := &models.Account{
		DocumentNumber: request.DocumentNumber,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	account, err := s.storage.CreateAccount(ctx, account)
	if err != nil {
		logger.GetLogger().Error("account creation failed")
		return "", &serviceErrors.ServiceError{WrappedError: err, Code: "internal-error", Description: "Internal WrappedError while creating account"}
	}
	return account, nil
}

func (s *PismoServiceImpl) FetchAccount(ctx context.Context, accountID string) (interface{}, *serviceErrors.ServiceError) {
	account, err := s.storage.FetchAccount(ctx, accountID)
	if err != nil {
		return nil, err
	}
	return account, nil
}

func (s *PismoServiceImpl) CreateTransaction(ctx context.Context) (interface{}, *serviceErrors.ServiceError) {
	return "new transaction created", nil
}
