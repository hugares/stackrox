// Code generated by MockGen. DO NOT EDIT.
// Source: aggregator.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	storage "github.com/stackrox/rox/generated/storage"
	reflect "reflect"
)

// MockNetworkConnsAggregator is a mock of NetworkConnsAggregator interface
type MockNetworkConnsAggregator struct {
	ctrl     *gomock.Controller
	recorder *MockNetworkConnsAggregatorMockRecorder
}

// MockNetworkConnsAggregatorMockRecorder is the mock recorder for MockNetworkConnsAggregator
type MockNetworkConnsAggregatorMockRecorder struct {
	mock *MockNetworkConnsAggregator
}

// NewMockNetworkConnsAggregator creates a new mock instance
func NewMockNetworkConnsAggregator(ctrl *gomock.Controller) *MockNetworkConnsAggregator {
	mock := &MockNetworkConnsAggregator{ctrl: ctrl}
	mock.recorder = &MockNetworkConnsAggregatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockNetworkConnsAggregator) EXPECT() *MockNetworkConnsAggregatorMockRecorder {
	return m.recorder
}

// Aggregate mocks base method
func (m *MockNetworkConnsAggregator) Aggregate(conns []*storage.NetworkFlow) []*storage.NetworkFlow {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Aggregate", conns)
	ret0, _ := ret[0].([]*storage.NetworkFlow)
	return ret0
}

// Aggregate indicates an expected call of Aggregate
func (mr *MockNetworkConnsAggregatorMockRecorder) Aggregate(conns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Aggregate", reflect.TypeOf((*MockNetworkConnsAggregator)(nil).Aggregate), conns)
}
