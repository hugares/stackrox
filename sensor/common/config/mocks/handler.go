// Code generated by MockGen. DO NOT EDIT.
// Source: handler.go
//
// Generated by this command:
//
//	mockgen -package mocks -destination mocks/handler.go -source handler.go
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	central "github.com/stackrox/rox/generated/internalapi/central"
	storage "github.com/stackrox/rox/generated/storage"
	centralsensor "github.com/stackrox/rox/pkg/centralsensor"
	common "github.com/stackrox/rox/sensor/common"
	message "github.com/stackrox/rox/sensor/common/message"
	gomock "go.uber.org/mock/gomock"
)

// MockHandler is a mock of Handler interface.
type MockHandler struct {
	ctrl     *gomock.Controller
	recorder *MockHandlerMockRecorder
}

// MockHandlerMockRecorder is the mock recorder for MockHandler.
type MockHandlerMockRecorder struct {
	mock *MockHandler
}

// NewMockHandler creates a new mock instance.
func NewMockHandler(ctrl *gomock.Controller) *MockHandler {
	mock := &MockHandler{ctrl: ctrl}
	mock.recorder = &MockHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHandler) EXPECT() *MockHandlerMockRecorder {
	return m.recorder
}

// Capabilities mocks base method.
func (m *MockHandler) Capabilities() []centralsensor.SensorCapability {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Capabilities")
	ret0, _ := ret[0].([]centralsensor.SensorCapability)
	return ret0
}

// Capabilities indicates an expected call of Capabilities.
func (mr *MockHandlerMockRecorder) Capabilities() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Capabilities", reflect.TypeOf((*MockHandler)(nil).Capabilities))
}

// GetConfig mocks base method.
func (m *MockHandler) GetConfig() *storage.DynamicClusterConfig {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConfig")
	ret0, _ := ret[0].(*storage.DynamicClusterConfig)
	return ret0
}

// GetConfig indicates an expected call of GetConfig.
func (mr *MockHandlerMockRecorder) GetConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConfig", reflect.TypeOf((*MockHandler)(nil).GetConfig))
}

// GetDeploymentIdentification mocks base method.
func (m *MockHandler) GetDeploymentIdentification() *storage.SensorDeploymentIdentification {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeploymentIdentification")
	ret0, _ := ret[0].(*storage.SensorDeploymentIdentification)
	return ret0
}

// GetDeploymentIdentification indicates an expected call of GetDeploymentIdentification.
func (mr *MockHandlerMockRecorder) GetDeploymentIdentification() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeploymentIdentification", reflect.TypeOf((*MockHandler)(nil).GetDeploymentIdentification))
}

// GetHelmManagedConfig mocks base method.
func (m *MockHandler) GetHelmManagedConfig() *central.HelmManagedConfigInit {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHelmManagedConfig")
	ret0, _ := ret[0].(*central.HelmManagedConfigInit)
	return ret0
}

// GetHelmManagedConfig indicates an expected call of GetHelmManagedConfig.
func (mr *MockHandlerMockRecorder) GetHelmManagedConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHelmManagedConfig", reflect.TypeOf((*MockHandler)(nil).GetHelmManagedConfig))
}

// Notify mocks base method.
func (m *MockHandler) Notify(e common.SensorComponentEvent) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Notify", e)
}

// Notify indicates an expected call of Notify.
func (mr *MockHandlerMockRecorder) Notify(e any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Notify", reflect.TypeOf((*MockHandler)(nil).Notify), e)
}

// ProcessMessage mocks base method.
func (m *MockHandler) ProcessMessage(msg *central.MsgToSensor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessMessage", msg)
	ret0, _ := ret[0].(error)
	return ret0
}

// ProcessMessage indicates an expected call of ProcessMessage.
func (mr *MockHandlerMockRecorder) ProcessMessage(msg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessMessage", reflect.TypeOf((*MockHandler)(nil).ProcessMessage), msg)
}

// ResponsesC mocks base method.
func (m *MockHandler) ResponsesC() <-chan *message.ExpiringMessage {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResponsesC")
	ret0, _ := ret[0].(<-chan *message.ExpiringMessage)
	return ret0
}

// ResponsesC indicates an expected call of ResponsesC.
func (mr *MockHandlerMockRecorder) ResponsesC() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResponsesC", reflect.TypeOf((*MockHandler)(nil).ResponsesC))
}

// Start mocks base method.
func (m *MockHandler) Start() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start")
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start.
func (mr *MockHandlerMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockHandler)(nil).Start))
}

// Stop mocks base method.
func (m *MockHandler) Stop(err error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Stop", err)
}

// Stop indicates an expected call of Stop.
func (mr *MockHandlerMockRecorder) Stop(err any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockHandler)(nil).Stop), err)
}
