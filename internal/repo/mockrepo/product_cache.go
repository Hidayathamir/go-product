// Code generated by MockGen. DO NOT EDIT.
// Source: product_cache.go
//
// Generated by this command:
//
//	mockgen -source=product_cache.go -destination=mockrepo/product_cache.go -package=mockrepo
//

// Package mockrepo is a generated GoMock package.
package mockrepo

import (
	context "context"
	reflect "reflect"
	time "time"

	goproduct "github.com/Hidayathamir/go-product/pkg/goproduct"
	gomock "go.uber.org/mock/gomock"
)

// MockIProductCache is a mock of IProductCache interface.
type MockIProductCache struct {
	ctrl     *gomock.Controller
	recorder *MockIProductCacheMockRecorder
}

// MockIProductCacheMockRecorder is the mock recorder for MockIProductCache.
type MockIProductCacheMockRecorder struct {
	mock *MockIProductCache
}

// NewMockIProductCache creates a new mock instance.
func NewMockIProductCache(ctrl *gomock.Controller) *MockIProductCache {
	mock := &MockIProductCache{ctrl: ctrl}
	mock.recorder = &MockIProductCacheMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIProductCache) EXPECT() *MockIProductCacheMockRecorder {
	return m.recorder
}

// GetDetailByID mocks base method.
func (m *MockIProductCache) GetDetailByID(ctx context.Context, ID int64) (goproduct.ResProductDetail, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDetailByID", ctx, ID)
	ret0, _ := ret[0].(goproduct.ResProductDetail)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDetailByID indicates an expected call of GetDetailByID.
func (mr *MockIProductCacheMockRecorder) GetDetailByID(ctx, ID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDetailByID", reflect.TypeOf((*MockIProductCache)(nil).GetDetailByID), ctx, ID)
}

// GetDetailBySKU mocks base method.
func (m *MockIProductCache) GetDetailBySKU(ctx context.Context, SKU string) (goproduct.ResProductDetail, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDetailBySKU", ctx, SKU)
	ret0, _ := ret[0].(goproduct.ResProductDetail)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDetailBySKU indicates an expected call of GetDetailBySKU.
func (mr *MockIProductCacheMockRecorder) GetDetailBySKU(ctx, SKU any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDetailBySKU", reflect.TypeOf((*MockIProductCache)(nil).GetDetailBySKU), ctx, SKU)
}

// GetDetailBySlug mocks base method.
func (m *MockIProductCache) GetDetailBySlug(ctx context.Context, slug string) (goproduct.ResProductDetail, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDetailBySlug", ctx, slug)
	ret0, _ := ret[0].(goproduct.ResProductDetail)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDetailBySlug indicates an expected call of GetDetailBySlug.
func (mr *MockIProductCacheMockRecorder) GetDetailBySlug(ctx, slug any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDetailBySlug", reflect.TypeOf((*MockIProductCache)(nil).GetDetailBySlug), ctx, slug)
}

// SetDetailByID mocks base method.
func (m *MockIProductCache) SetDetailByID(ctx context.Context, data goproduct.ResProductDetail, expire time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetDetailByID", ctx, data, expire)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetDetailByID indicates an expected call of SetDetailByID.
func (mr *MockIProductCacheMockRecorder) SetDetailByID(ctx, data, expire any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDetailByID", reflect.TypeOf((*MockIProductCache)(nil).SetDetailByID), ctx, data, expire)
}

// SetDetailBySKU mocks base method.
func (m *MockIProductCache) SetDetailBySKU(ctx context.Context, data goproduct.ResProductDetail, expire time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetDetailBySKU", ctx, data, expire)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetDetailBySKU indicates an expected call of SetDetailBySKU.
func (mr *MockIProductCacheMockRecorder) SetDetailBySKU(ctx, data, expire any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDetailBySKU", reflect.TypeOf((*MockIProductCache)(nil).SetDetailBySKU), ctx, data, expire)
}

// SetDetailBySlug mocks base method.
func (m *MockIProductCache) SetDetailBySlug(ctx context.Context, data goproduct.ResProductDetail, expire time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetDetailBySlug", ctx, data, expire)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetDetailBySlug indicates an expected call of SetDetailBySlug.
func (mr *MockIProductCacheMockRecorder) SetDetailBySlug(ctx, data, expire any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDetailBySlug", reflect.TypeOf((*MockIProductCache)(nil).SetDetailBySlug), ctx, data, expire)
}