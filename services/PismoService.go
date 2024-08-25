package services

import "context"

//go:generate mockgen -source=PismoService.go -destination=./mock_services/mock_pismo_service.go
type PismoService interface {
	Greet(ctx context.Context) string
	CreateAccount(ctx context.Context) string
	FetchAccount(ctx context.Context, accountID string) string
	CreateTransaction(ctx context.Context) string
}

type PismoServiceImpl struct{}

func (s *PismoServiceImpl) Greet(ctx context.Context) string {
	return "Greetings from Pismo-Test"
}

func (s *PismoServiceImpl) CreateAccount(ctx context.Context) string {
	return "Account has been created"
}

func (s *PismoServiceImpl) FetchAccount(ctx context.Context, accountID string) string {
	return "account fetched"
}

func (s *PismoServiceImpl) CreateTransaction(ctx context.Context) string {
	return "new transaction created"
}
