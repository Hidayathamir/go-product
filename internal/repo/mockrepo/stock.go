// Code generated by MockGen. DO NOT EDIT.
// Source: stock.go
//
// Generated by this command:
//
//	mockgen -source=stock.go -destination=mockrepo/stock.go -package=mockrepo
//

// Package mockrepo is a generated GoMock package.
package mockrepo

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockIStock is a mock of IStock interface.
type MockIStock struct {
	ctrl     *gomock.Controller
	recorder *MockIStockMockRecorder
}

// MockIStockMockRecorder is the mock recorder for MockIStock.
type MockIStockMockRecorder struct {
	mock *MockIStock
}

// NewMockIStock creates a new mock instance.
func NewMockIStock(ctrl *gomock.Controller) *MockIStock {
	mock := &MockIStock{ctrl: ctrl}
	mock.recorder = &MockIStockMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIStock) EXPECT() *MockIStockMockRecorder {
	return m.recorder
}

// DecrementStock mocks base method.
func (m *MockIStock) DecrementStock(ctx context.Context, productID int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DecrementStock", ctx, productID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DecrementStock indicates an expected call of DecrementStock.
func (mr *MockIStockMockRecorder) DecrementStock(ctx, productID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DecrementStock", reflect.TypeOf((*MockIStock)(nil).DecrementStock), ctx, productID)
}

// IncrementStock mocks base method.
func (m *MockIStock) IncrementStock(ctx context.Context, productID int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IncrementStock", ctx, productID)
	ret0, _ := ret[0].(error)
	return ret0
}

// IncrementStock indicates an expected call of IncrementStock.
func (mr *MockIStockMockRecorder) IncrementStock(ctx, productID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IncrementStock", reflect.TypeOf((*MockIStock)(nil).IncrementStock), ctx, productID)
}