// Code generated by MockGen. DO NOT EDIT.
// Source: indexer.go
//
// Generated by this command:
//
//	mockgen -package mocks -destination mocks/indexer.go -source indexer.go
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	claircore "github.com/quay/claircore"
	indexer "github.com/stackrox/rox/scanner/indexer"
	gomock "go.uber.org/mock/gomock"
)

// MockReportGetter is a mock of ReportGetter interface.
type MockReportGetter struct {
	ctrl     *gomock.Controller
	recorder *MockReportGetterMockRecorder
}

// MockReportGetterMockRecorder is the mock recorder for MockReportGetter.
type MockReportGetterMockRecorder struct {
	mock *MockReportGetter
}

// NewMockReportGetter creates a new mock instance.
func NewMockReportGetter(ctrl *gomock.Controller) *MockReportGetter {
	mock := &MockReportGetter{ctrl: ctrl}
	mock.recorder = &MockReportGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReportGetter) EXPECT() *MockReportGetterMockRecorder {
	return m.recorder
}

// GetIndexReport mocks base method.
func (m *MockReportGetter) GetIndexReport(arg0 context.Context, arg1 string) (*claircore.IndexReport, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIndexReport", arg0, arg1)
	ret0, _ := ret[0].(*claircore.IndexReport)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetIndexReport indicates an expected call of GetIndexReport.
func (mr *MockReportGetterMockRecorder) GetIndexReport(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIndexReport", reflect.TypeOf((*MockReportGetter)(nil).GetIndexReport), arg0, arg1)
}

// MockIndexer is a mock of Indexer interface.
type MockIndexer struct {
	ctrl     *gomock.Controller
	recorder *MockIndexerMockRecorder
}

// MockIndexerMockRecorder is the mock recorder for MockIndexer.
type MockIndexerMockRecorder struct {
	mock *MockIndexer
}

// NewMockIndexer creates a new mock instance.
func NewMockIndexer(ctrl *gomock.Controller) *MockIndexer {
	mock := &MockIndexer{ctrl: ctrl}
	mock.recorder = &MockIndexerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIndexer) EXPECT() *MockIndexerMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockIndexer) Close(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockIndexerMockRecorder) Close(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockIndexer)(nil).Close), arg0)
}

// GetIndexReport mocks base method.
func (m *MockIndexer) GetIndexReport(arg0 context.Context, arg1 string) (*claircore.IndexReport, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIndexReport", arg0, arg1)
	ret0, _ := ret[0].(*claircore.IndexReport)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetIndexReport indicates an expected call of GetIndexReport.
func (mr *MockIndexerMockRecorder) GetIndexReport(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIndexReport", reflect.TypeOf((*MockIndexer)(nil).GetIndexReport), arg0, arg1)
}

// IndexContainerImage mocks base method.
func (m *MockIndexer) IndexContainerImage(arg0 context.Context, arg1, arg2 string, arg3 ...indexer.Option) (*claircore.IndexReport, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "IndexContainerImage", varargs...)
	ret0, _ := ret[0].(*claircore.IndexReport)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IndexContainerImage indicates an expected call of IndexContainerImage.
func (mr *MockIndexerMockRecorder) IndexContainerImage(arg0, arg1, arg2 any, arg3 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IndexContainerImage", reflect.TypeOf((*MockIndexer)(nil).IndexContainerImage), varargs...)
}
