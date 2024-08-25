package services_test

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/sudipidus/pismo-test/logger"
	"github.com/sudipidus/pismo-test/services"
	"github.com/sudipidus/pismo-test/storage/mock_storage"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestPismoServiceImpl_CreateAccount(t *testing.T) {
	t.Run("should return good for successful account creation", func(t *testing.T) {
		logger.InitLogger()
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockStorage := mock_storage.NewMockStorage(ctrl)
		mockStorage.EXPECT().CreateAccount(gomock.Any()).Return(errors.New("error creating account"))

		service := services.NewPismoService(mockStorage)
		response, err := service.CreateAccount(context.Background())
		assert.Empty(t, response)
		assert.Equal(t, errors.New("account creation failed"), err)

	})
}
