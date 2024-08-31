// Code generated by MockGen. DO NOT EDIT.
// Source: PismoService.go

// Package mock_services is a generated GoMock package.
package mock_services

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	errors "github.com/sudipidus/pismo-test/errors"
	models "github.com/sudipidus/pismo-test/models"
	services "github.com/sudipidus/pismo-test/services"
)

// MockPismoService is a mock of PismoService interface.
type MockPismoService struct {
	ctrl     *gomock.Controller
	recorder *MockPismoServiceMockRecorder
}

// MockPismoServiceMockRecorder is the mock recorder for MockPismoService.
type MockPismoServiceMockRecorder struct {
	mock *MockPismoService
}

// NewMockPismoService creates a new mock instance.
func NewMockPismoService(ctrl *gomock.Controller) *MockPismoService {
	mock := &MockPismoService{ctrl: ctrl}
	mock.recorder = &MockPismoServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPismoService) EXPECT() *MockPismoServiceMockRecorder {
	return m.recorder
}

// CreateAccount mocks base method.
func (m *MockPismoService) CreateAccount(ctx context.Context, request services.CreateAccountRequest) (*models.Account, *errors.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAccount", ctx, request)
	ret0, _ := ret[0].(*models.Account)
	ret1, _ := ret[1].(*errors.Error)
	return ret0, ret1
}

// CreateAccount indicates an expected call of CreateAccount.
func (mr *MockPismoServiceMockRecorder) CreateAccount(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccount", reflect.TypeOf((*MockPismoService)(nil).CreateAccount), ctx, request)
}

// CreateTransaction mocks base method.
func (m *MockPismoService) CreateTransaction(ctx context.Context, request services.CreateTransactionRequest) (*models.Transaction, *errors.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTransaction", ctx, request)
	ret0, _ := ret[0].(*models.Transaction)
	ret1, _ := ret[1].(*errors.Error)
	return ret0, ret1
}

// CreateTransaction indicates an expected call of CreateTransaction.
func (mr *MockPismoServiceMockRecorder) CreateTransaction(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTransaction", reflect.TypeOf((*MockPismoService)(nil).CreateTransaction), ctx, request)
}

// FetchAccount mocks base method.
func (m *MockPismoService) FetchAccount(ctx context.Context, accountID string) (*models.Account, *errors.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchAccount", ctx, accountID)
	ret0, _ := ret[0].(*models.Account)
	ret1, _ := ret[1].(*errors.Error)
	return ret0, ret1
}

// FetchAccount indicates an expected call of FetchAccount.
func (mr *MockPismoServiceMockRecorder) FetchAccount(ctx, accountID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchAccount", reflect.TypeOf((*MockPismoService)(nil).FetchAccount), ctx, accountID)
}

// Greet mocks base method.
func (m *MockPismoService) Greet(ctx context.Context) (interface{}, *errors.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Greet", ctx)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(*errors.Error)
	return ret0, ret1
}

// Greet indicates an expected call of Greet.
func (mr *MockPismoServiceMockRecorder) Greet(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Greet", reflect.TypeOf((*MockPismoService)(nil).Greet), ctx)
}
