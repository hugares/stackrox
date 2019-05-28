// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/stackrox/rox/central/serviceaccount/internal/store (interfaces: Store)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	storage "github.com/stackrox/rox/generated/storage"
	reflect "reflect"
)

// MockStore is a mock of Store interface
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// DeleteServiceAccount mocks base method
func (m *MockStore) DeleteServiceAccount(arg0 string) error {
	ret := m.ctrl.Call(m, "DeleteServiceAccount", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteServiceAccount indicates an expected call of DeleteServiceAccount
func (mr *MockStoreMockRecorder) DeleteServiceAccount(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteServiceAccount", reflect.TypeOf((*MockStore)(nil).DeleteServiceAccount), arg0)
}

// GetServiceAccount mocks base method
func (m *MockStore) GetServiceAccount(arg0 string) (*storage.ServiceAccount, bool, error) {
	ret := m.ctrl.Call(m, "GetServiceAccount", arg0)
	ret0, _ := ret[0].(*storage.ServiceAccount)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetServiceAccount indicates an expected call of GetServiceAccount
func (mr *MockStoreMockRecorder) GetServiceAccount(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetServiceAccount", reflect.TypeOf((*MockStore)(nil).GetServiceAccount), arg0)
}

// GetServiceAccounts mocks base method
func (m *MockStore) GetServiceAccounts(arg0 []string) ([]*storage.ServiceAccount, []int, error) {
	ret := m.ctrl.Call(m, "GetServiceAccounts", arg0)
	ret0, _ := ret[0].([]*storage.ServiceAccount)
	ret1, _ := ret[1].([]int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetServiceAccounts indicates an expected call of GetServiceAccounts
func (mr *MockStoreMockRecorder) GetServiceAccounts(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetServiceAccounts", reflect.TypeOf((*MockStore)(nil).GetServiceAccounts), arg0)
}

// ListServiceAccounts mocks base method
func (m *MockStore) ListServiceAccounts() ([]*storage.ServiceAccount, error) {
	ret := m.ctrl.Call(m, "ListServiceAccounts")
	ret0, _ := ret[0].([]*storage.ServiceAccount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListServiceAccounts indicates an expected call of ListServiceAccounts
func (mr *MockStoreMockRecorder) ListServiceAccounts() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListServiceAccounts", reflect.TypeOf((*MockStore)(nil).ListServiceAccounts))
}

// UpsertServiceAccount mocks base method
func (m *MockStore) UpsertServiceAccount(arg0 *storage.ServiceAccount) error {
	ret := m.ctrl.Call(m, "UpsertServiceAccount", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertServiceAccount indicates an expected call of UpsertServiceAccount
func (mr *MockStoreMockRecorder) UpsertServiceAccount(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertServiceAccount", reflect.TypeOf((*MockStore)(nil).UpsertServiceAccount), arg0)
}
