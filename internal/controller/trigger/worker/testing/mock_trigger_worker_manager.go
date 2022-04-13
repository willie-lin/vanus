// Code generated by MockGen. DO NOT EDIT.
// Source: trigger_worker_manager.go

// Package testing is a generated GoMock package.
package testing

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	info "github.com/linkall-labs/vanus/internal/controller/trigger/info"
)

// MockManager is a mock of Manager interface.
type MockManager struct {
	ctrl     *gomock.Controller
	recorder *MockManagerMockRecorder
}

// MockManagerMockRecorder is the mock recorder for MockManager.
type MockManagerMockRecorder struct {
	mock *MockManager
}

// NewMockManager creates a new mock instance.
func NewMockManager(ctrl *gomock.Controller) *MockManager {
	mock := &MockManager{ctrl: ctrl}
	mock.recorder = &MockManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockManager) EXPECT() *MockManagerMockRecorder {
	return m.recorder
}

// AddTriggerWorker mocks base method.
func (m *MockManager) AddTriggerWorker(ctx context.Context, addr string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddTriggerWorker", ctx, addr)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddTriggerWorker indicates an expected call of AddTriggerWorker.
func (mr *MockManagerMockRecorder) AddTriggerWorker(ctx, addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTriggerWorker", reflect.TypeOf((*MockManager)(nil).AddTriggerWorker), ctx, addr)
}

// AssignSubscription mocks base method.
func (m *MockManager) AssignSubscription(ctx context.Context, twInfo info.TriggerWorkerInfo, subId string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AssignSubscription", ctx, twInfo, subId)
	ret0, _ := ret[0].(error)
	return ret0
}

// AssignSubscription indicates an expected call of AssignSubscription.
func (mr *MockManagerMockRecorder) AssignSubscription(ctx, twInfo, subId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssignSubscription", reflect.TypeOf((*MockManager)(nil).AssignSubscription), ctx, twInfo, subId)
}

// GetRunningTriggerWorker mocks base method.
func (m *MockManager) GetRunningTriggerWorker() []info.TriggerWorkerInfo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRunningTriggerWorker")
	ret0, _ := ret[0].([]info.TriggerWorkerInfo)
	return ret0
}

// GetRunningTriggerWorker indicates an expected call of GetRunningTriggerWorker.
func (mr *MockManagerMockRecorder) GetRunningTriggerWorker() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRunningTriggerWorker", reflect.TypeOf((*MockManager)(nil).GetRunningTriggerWorker))
}

// Init mocks base method.
func (m *MockManager) Init(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Init", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Init indicates an expected call of Init.
func (mr *MockManagerMockRecorder) Init(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockManager)(nil).Init), ctx)
}

// RemoveTriggerWorker mocks base method.
func (m *MockManager) RemoveTriggerWorker(ctx context.Context, addr string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RemoveTriggerWorker", ctx, addr)
}

// RemoveTriggerWorker indicates an expected call of RemoveTriggerWorker.
func (mr *MockManagerMockRecorder) RemoveTriggerWorker(ctx, addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveTriggerWorker", reflect.TypeOf((*MockManager)(nil).RemoveTriggerWorker), ctx, addr)
}

// Start mocks base method.
func (m *MockManager) Start() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Start")
}

// Start indicates an expected call of Start.
func (mr *MockManagerMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockManager)(nil).Start))
}

// Stop mocks base method.
func (m *MockManager) Stop() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Stop")
}

// Stop indicates an expected call of Stop.
func (mr *MockManagerMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockManager)(nil).Stop))
}

// UnAssignSubscription mocks base method.
func (m *MockManager) UnAssignSubscription(ctx context.Context, addr, subId string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnAssignSubscription", ctx, addr, subId)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnAssignSubscription indicates an expected call of UnAssignSubscription.
func (mr *MockManagerMockRecorder) UnAssignSubscription(ctx, addr, subId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnAssignSubscription", reflect.TypeOf((*MockManager)(nil).UnAssignSubscription), ctx, addr, subId)
}

// UpdateTriggerWorkerInfo mocks base method.
func (m *MockManager) UpdateTriggerWorkerInfo(ctx context.Context, addr string, subIds map[string]struct{}) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTriggerWorkerInfo", ctx, addr, subIds)
	ret0, _ := ret[0].(bool)
	return ret0
}

// UpdateTriggerWorkerInfo indicates an expected call of UpdateTriggerWorkerInfo.
func (mr *MockManagerMockRecorder) UpdateTriggerWorkerInfo(ctx, addr, subIds interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTriggerWorkerInfo", reflect.TypeOf((*MockManager)(nil).UpdateTriggerWorkerInfo), ctx, addr, subIds)
}
