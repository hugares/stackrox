// Code generated by MockGen. DO NOT EDIT.
// Source: store.go

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

// ListAlert mocks base method
func (m *MockStore) ListAlert(id string) (*storage.ListAlert, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAlert", id)
	ret0, _ := ret[0].(*storage.ListAlert)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListAlert indicates an expected call of ListAlert
func (mr *MockStoreMockRecorder) ListAlert(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAlert", reflect.TypeOf((*MockStore)(nil).ListAlert), id)
}

// ListAlerts mocks base method
func (m *MockStore) ListAlerts() ([]*storage.ListAlert, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAlerts")
	ret0, _ := ret[0].([]*storage.ListAlert)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAlerts indicates an expected call of ListAlerts
func (mr *MockStoreMockRecorder) ListAlerts() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAlerts", reflect.TypeOf((*MockStore)(nil).ListAlerts))
}

// GetListAlerts mocks base method
func (m *MockStore) GetListAlerts(arg0 []string) ([]*storage.ListAlert, []int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetListAlerts", arg0)
	ret0, _ := ret[0].([]*storage.ListAlert)
	ret1, _ := ret[1].([]int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetListAlerts indicates an expected call of GetListAlerts
func (mr *MockStoreMockRecorder) GetListAlerts(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetListAlerts", reflect.TypeOf((*MockStore)(nil).GetListAlerts), arg0)
}

// GetAlertIDs mocks base method
func (m *MockStore) GetAlertIDs() ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAlertIDs")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAlertIDs indicates an expected call of GetAlertIDs
func (mr *MockStoreMockRecorder) GetAlertIDs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAlertIDs", reflect.TypeOf((*MockStore)(nil).GetAlertIDs))
}

// GetAlert mocks base method
func (m *MockStore) GetAlert(id string) (*storage.Alert, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAlert", id)
	ret0, _ := ret[0].(*storage.Alert)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetAlert indicates an expected call of GetAlert
func (mr *MockStoreMockRecorder) GetAlert(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAlert", reflect.TypeOf((*MockStore)(nil).GetAlert), id)
}

// GetAlerts mocks base method
func (m *MockStore) GetAlerts(ids []string) ([]*storage.Alert, []int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAlerts", ids)
	ret0, _ := ret[0].([]*storage.Alert)
	ret1, _ := ret[1].([]int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetAlerts indicates an expected call of GetAlerts
func (mr *MockStoreMockRecorder) GetAlerts(ids interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAlerts", reflect.TypeOf((*MockStore)(nil).GetAlerts), ids)
}

// AddAlert mocks base method
func (m *MockStore) AddAlert(alert *storage.Alert) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddAlert", alert)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddAlert indicates an expected call of AddAlert
func (mr *MockStoreMockRecorder) AddAlert(alert interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAlert", reflect.TypeOf((*MockStore)(nil).AddAlert), alert)
}

// UpdateAlert mocks base method
func (m *MockStore) UpdateAlert(alert *storage.Alert) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAlert", alert)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAlert indicates an expected call of UpdateAlert
func (mr *MockStoreMockRecorder) UpdateAlert(alert interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAlert", reflect.TypeOf((*MockStore)(nil).UpdateAlert), alert)
}

// DeleteAlert mocks base method
func (m *MockStore) DeleteAlert(id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAlert", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAlert indicates an expected call of DeleteAlert
func (mr *MockStoreMockRecorder) DeleteAlert(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAlert", reflect.TypeOf((*MockStore)(nil).DeleteAlert), id)
}

// DeleteAlerts mocks base method
func (m *MockStore) DeleteAlerts(ids ...string) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range ids {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteAlerts", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAlerts indicates an expected call of DeleteAlerts
func (mr *MockStoreMockRecorder) DeleteAlerts(ids ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAlerts", reflect.TypeOf((*MockStore)(nil).DeleteAlerts), ids...)
}

// GetTxnCount mocks base method
func (m *MockStore) GetTxnCount() (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTxnCount")
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTxnCount indicates an expected call of GetTxnCount
func (mr *MockStoreMockRecorder) GetTxnCount() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTxnCount", reflect.TypeOf((*MockStore)(nil).GetTxnCount))
}

// IncTxnCount mocks base method
func (m *MockStore) IncTxnCount() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IncTxnCount")
	ret0, _ := ret[0].(error)
	return ret0
}

// IncTxnCount indicates an expected call of IncTxnCount
func (mr *MockStoreMockRecorder) IncTxnCount() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IncTxnCount", reflect.TypeOf((*MockStore)(nil).IncTxnCount))
}
