package services_test

import (
	"context"
	errors2 "errors"
	"github.com/stretchr/testify/assert"
	"github.com/sudipidus/pismo-test/errors"
	"github.com/sudipidus/pismo-test/logger"
	"github.com/sudipidus/pismo-test/models"
	"github.com/sudipidus/pismo-test/services"
	"github.com/sudipidus/pismo-test/storage/mock_storage"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestPismoServiceImpl_CreateAccount(t *testing.T) {
	t.Run("service should successfully create and return an account if the storage layer succeeds in creating the account", func(t *testing.T) {
		logger.InitLogger()
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		account := &models.Account{DocumentNumber: "1234567890"}

		mockStorage := mock_storage.NewMockStorage(ctrl)
		mockStorage.EXPECT().CreateAccount(context.Background(), &accountMatcher{expected: account}).
			Return(&models.Account{DocumentNumber: "1234567890"}, nil)

		service := services.NewPismoService(mockStorage)
		response, err := service.CreateAccount(context.Background(), services.CreateAccountRequest{
			DocumentNumber: "1234567890",
		})
		assert.Nil(t, err)
		assert.NotEmpty(t, response)
	})

	t.Run("service should return error if the storage layer errors out", func(t *testing.T) {
		logger.InitLogger()
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		account := &models.Account{DocumentNumber: "1234567890"}

		mockStorage := mock_storage.NewMockStorage(ctrl)
		mockStorage.EXPECT().CreateAccount(context.Background(), &accountMatcher{expected: account}).
			Return(nil, &errors.Error{
				Err:     errors2.New("failed to create entry"),
				Code:    500,
				Message: "failed to create entry",
			})

		service := services.NewPismoService(mockStorage)
		response, err := service.CreateAccount(context.Background(), services.CreateAccountRequest{
			DocumentNumber: "1234567890",
		})
		assert.Empty(t, response)
		assert.Equal(t, err.Error(), "failed to create entry")
		assert.Equal(t, err.Code, 500)
	})
}

func TestPismoServiceImpl_FetchAccount(t *testing.T) {
	t.Run("service should successfully return an account if the storage layer succeeds in returning the account", func(t *testing.T) {
		logger.InitLogger()
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		accountID, documentNumber := "1", "1234567890"

		mockStorage := mock_storage.NewMockStorage(ctrl)
		mockStorage.EXPECT().FetchAccount(context.Background(), accountID).
			Return(&models.Account{DocumentNumber: "1234567890", ID: 1}, nil)

		service := services.NewPismoService(mockStorage)
		response, err := service.FetchAccount(context.Background(), accountID)
		assert.Nil(t, err)
		assert.NotEmpty(t, response)
		assert.Equal(t, documentNumber, response.DocumentNumber)
		assert.Equal(t, 1, response.ID)
	})

	t.Run("service should return error if the storage layer errors out while fetching account", func(t *testing.T) {
		logger.InitLogger()
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		accountID, _ := "1", "1234567890"

		mockStorage := mock_storage.NewMockStorage(ctrl)
		mockStorage.EXPECT().FetchAccount(context.Background(), accountID).
			Return(nil, &errors.Error{
				Err:     errors2.New("failed to fetch account"),
				Code:    500,
				Message: "failed to fetch account",
			})

		service := services.NewPismoService(mockStorage)
		response, err := service.FetchAccount(context.Background(), accountID)
		assert.Empty(t, response)
		assert.Equal(t, err.Error(), "failed to fetch account")
	})
}

type accountMatcher struct {
	expected *models.Account
}

func (m *accountMatcher) Matches(x interface{}) bool {
	account, ok := x.(*models.Account)
	if !ok {
		return false
	}
	return account.DocumentNumber == m.expected.DocumentNumber
}

func (m *accountMatcher) String() string {
	return "matches account with specified DocumentNumber"
}
