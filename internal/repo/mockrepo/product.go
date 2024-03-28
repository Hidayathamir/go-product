// Code generated by MockGen. DO NOT EDIT.
// Source: product.go
//
// Generated by this command:
//
//	mockgen -source=product.go -destination=mockrepo/product.go -package=mockrepo
//

// Package mockrepo is a generated GoMock package.
package mockrepo

import (
	context "context"
	reflect "reflect"

	goproduct "github.com/Hidayathamir/go-product/pkg/goproduct"
	gomock "go.uber.org/mock/gomock"
)

// MockIProduct is a mock of IProduct interface.
type MockIProduct struct {
	ctrl     *gomock.Controller
	recorder *MockIProductMockRecorder
}

// MockIProductMockRecorder is the mock recorder for MockIProduct.
type MockIProductMockRecorder struct {
	mock *MockIProduct
}

// NewMockIProduct creates a new mock instance.
func NewMockIProduct(ctrl *gomock.Controller) *MockIProduct {
	mock := &MockIProduct{ctrl: ctrl}
	mock.recorder = &MockIProductMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIProduct) EXPECT() *MockIProductMockRecorder {
	return m.recorder
}

// GetDetailByID mocks base method.
func (m *MockIProduct) GetDetailByID(ctx context.Context, ID int64) (goproduct.ResProductDetail, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDetailByID", ctx, ID)
	ret0, _ := ret[0].(goproduct.ResProductDetail)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDetailByID indicates an expected call of GetDetailByID.
func (mr *MockIProductMockRecorder) GetDetailByID(ctx, ID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDetailByID", reflect.TypeOf((*MockIProduct)(nil).GetDetailByID), ctx, ID)
}

// GetDetailBySKU mocks base method.
func (m *MockIProduct) GetDetailBySKU(ctx context.Context, SKU string) (goproduct.ResProductDetail, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDetailBySKU", ctx, SKU)
	ret0, _ := ret[0].(goproduct.ResProductDetail)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDetailBySKU indicates an expected call of GetDetailBySKU.
func (mr *MockIProductMockRecorder) GetDetailBySKU(ctx, SKU any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDetailBySKU", reflect.TypeOf((*MockIProduct)(nil).GetDetailBySKU), ctx, SKU)
}

// GetDetailBySlug mocks base method.
func (m *MockIProduct) GetDetailBySlug(ctx context.Context, slug string) (goproduct.ResProductDetail, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDetailBySlug", ctx, slug)
	ret0, _ := ret[0].(goproduct.ResProductDetail)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDetailBySlug indicates an expected call of GetDetailBySlug.
func (mr *MockIProductMockRecorder) GetDetailBySlug(ctx, slug any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDetailBySlug", reflect.TypeOf((*MockIProduct)(nil).GetDetailBySlug), ctx, slug)
}

// Search mocks base method.
func (m *MockIProduct) Search(ctx context.Context, keyword string) (goproduct.ResProductSearch, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Search", ctx, keyword)
	ret0, _ := ret[0].(goproduct.ResProductSearch)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Search indicates an expected call of Search.
func (mr *MockIProductMockRecorder) Search(ctx, keyword any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockIProduct)(nil).Search), ctx, keyword)
}