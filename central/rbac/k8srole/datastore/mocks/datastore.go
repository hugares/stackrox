// Code generated by MockGen. DO NOT EDIT.
// Source: datastore.go
//
// Generated by this command:
//
//	mockgen -package mocks -destination mocks/datastore.go -source datastore.go
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	v1 "github.com/stackrox/rox/generated/api/v1"
	storage "github.com/stackrox/rox/generated/storage"
	search "github.com/stackrox/rox/pkg/search"
	gomock "go.uber.org/mock/gomock"
)

// MockDataStore is a mock of DataStore interface.
type MockDataStore struct {
	ctrl     *gomock.Controller
	recorder *MockDataStoreMockRecorder
}

// MockDataStoreMockRecorder is the mock recorder for MockDataStore.
type MockDataStoreMockRecorder struct {
	mock *MockDataStore
}

// NewMockDataStore creates a new mock instance.
func NewMockDataStore(ctrl *gomock.Controller) *MockDataStore {
	mock := &MockDataStore{ctrl: ctrl}
	mock.recorder = &MockDataStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDataStore) EXPECT() *MockDataStoreMockRecorder {
	return m.recorder
}

// Count mocks base method.
func (m *MockDataStore) Count(ctx context.Context, q *v1.Query) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", ctx, q)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockDataStoreMockRecorder) Count(ctx, q any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockDataStore)(nil).Count), ctx, q)
}

// GetRole mocks base method.
func (m *MockDataStore) GetRole(ctx context.Context, id string) (*storage.K8SRole, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRole", ctx, id)
	ret0, _ := ret[0].(*storage.K8SRole)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetRole indicates an expected call of GetRole.
func (mr *MockDataStoreMockRecorder) GetRole(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRole", reflect.TypeOf((*MockDataStore)(nil).GetRole), ctx, id)
}

// RemoveRole mocks base method.
func (m *MockDataStore) RemoveRole(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveRole", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveRole indicates an expected call of RemoveRole.
func (mr *MockDataStoreMockRecorder) RemoveRole(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveRole", reflect.TypeOf((*MockDataStore)(nil).RemoveRole), ctx, id)
}

// Search mocks base method.
func (m *MockDataStore) Search(ctx context.Context, q *v1.Query) ([]search.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Search", ctx, q)
	ret0, _ := ret[0].([]search.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Search indicates an expected call of Search.
func (mr *MockDataStoreMockRecorder) Search(ctx, q any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockDataStore)(nil).Search), ctx, q)
}

// SearchRawRoles mocks base method.
func (m *MockDataStore) SearchRawRoles(ctx context.Context, q *v1.Query) ([]*storage.K8SRole, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchRawRoles", ctx, q)
	ret0, _ := ret[0].([]*storage.K8SRole)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchRawRoles indicates an expected call of SearchRawRoles.
func (mr *MockDataStoreMockRecorder) SearchRawRoles(ctx, q any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchRawRoles", reflect.TypeOf((*MockDataStore)(nil).SearchRawRoles), ctx, q)
}

// SearchRoles mocks base method.
func (m *MockDataStore) SearchRoles(ctx context.Context, q *v1.Query) ([]*v1.SearchResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchRoles", ctx, q)
	ret0, _ := ret[0].([]*v1.SearchResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchRoles indicates an expected call of SearchRoles.
func (mr *MockDataStoreMockRecorder) SearchRoles(ctx, q any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchRoles", reflect.TypeOf((*MockDataStore)(nil).SearchRoles), ctx, q)
}

// UpsertRole mocks base method.
func (m *MockDataStore) UpsertRole(ctx context.Context, request *storage.K8SRole) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertRole", ctx, request)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertRole indicates an expected call of UpsertRole.
func (mr *MockDataStoreMockRecorder) UpsertRole(ctx, request any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertRole", reflect.TypeOf((*MockDataStore)(nil).UpsertRole), ctx, request)
}
