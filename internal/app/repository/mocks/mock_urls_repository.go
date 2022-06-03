// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/iamtonydev/url-shortener/internal/app/repository (interfaces: IUrlsRepository)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIUrlsRepository is a mock of IUrlsRepository interface.
type MockIUrlsRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIUrlsRepositoryMockRecorder
}

// MockIUrlsRepositoryMockRecorder is the mock recorder for MockIUrlsRepository.
type MockIUrlsRepositoryMockRecorder struct {
	mock *MockIUrlsRepository
}

// NewMockIUrlsRepository creates a new mock instance.
func NewMockIUrlsRepository(ctrl *gomock.Controller) *MockIUrlsRepository {
	mock := &MockIUrlsRepository{ctrl: ctrl}
	mock.recorder = &MockIUrlsRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIUrlsRepository) EXPECT() *MockIUrlsRepositoryMockRecorder {
	return m.recorder
}

// AddShortUrl mocks base method.
func (m *MockIUrlsRepository) AddShortUrl(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddShortUrl", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddShortUrl indicates an expected call of AddShortUrl.
func (mr *MockIUrlsRepositoryMockRecorder) AddShortUrl(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddShortUrl", reflect.TypeOf((*MockIUrlsRepository)(nil).AddShortUrl), arg0, arg1, arg2)
}

// GetUrl mocks base method.
func (m *MockIUrlsRepository) GetUrl(arg0 context.Context, arg1 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUrl", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUrl indicates an expected call of GetUrl.
func (mr *MockIUrlsRepositoryMockRecorder) GetUrl(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUrl", reflect.TypeOf((*MockIUrlsRepository)(nil).GetUrl), arg0, arg1)
}

// IsNotFoundError mocks base method.
func (m *MockIUrlsRepository) IsNotFoundError(arg0 error) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsNotFoundError", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsNotFoundError indicates an expected call of IsNotFoundError.
func (mr *MockIUrlsRepositoryMockRecorder) IsNotFoundError(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsNotFoundError", reflect.TypeOf((*MockIUrlsRepository)(nil).IsNotFoundError), arg0)
}

// IsShortUrlDuplicateError mocks base method.
func (m *MockIUrlsRepository) IsShortUrlDuplicateError(arg0 error) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsShortUrlDuplicateError", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsShortUrlDuplicateError indicates an expected call of IsShortUrlDuplicateError.
func (mr *MockIUrlsRepositoryMockRecorder) IsShortUrlDuplicateError(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsShortUrlDuplicateError", reflect.TypeOf((*MockIUrlsRepository)(nil).IsShortUrlDuplicateError), arg0)
}

// IsUrlDuplicateError mocks base method.
func (m *MockIUrlsRepository) IsUrlDuplicateError(arg0 error) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsUrlDuplicateError", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsUrlDuplicateError indicates an expected call of IsUrlDuplicateError.
func (mr *MockIUrlsRepositoryMockRecorder) IsUrlDuplicateError(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsUrlDuplicateError", reflect.TypeOf((*MockIUrlsRepository)(nil).IsUrlDuplicateError), arg0)
}
