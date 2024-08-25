package services

import "context"

//go:generate mockgen -source=PismoService.go -destination=./mock_services/mock_pismo_service.go
type PismoService interface {
	Greet(ctx context.Context) string
	CreateAccount(ctx context.Context) string
	FetchAccount(ctx context.Context) string
	CreateTransaction(ctx context.Context) string
}

type pismoServiceImpl struct{}

func (s *pismoServiceImpl) Greet(ctx context.Context) string {
	return "Greetings from Pismo-Test"
}

func (s *pismoServiceImpl) CreateAccount(ctx context.Context) string {
	return "Account has been created"
}

func (s *pismoServiceImpl) GetAccountByID(ctx context.Context, accountID string) string {
	return "account fetched"
}

func (s *pismoServiceImpl) CreateTransaction(ctx context.Context) string {
	return "new transaction created"
}
