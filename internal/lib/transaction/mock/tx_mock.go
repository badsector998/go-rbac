// Code generated by MockGen. DO NOT EDIT.
// Source: D:\tugas\Github\go-rbac\internal\lib\transaction\tx.go
//
// Generated by this command:
//
//	mockgen -source=D:\tugas\Github\go-rbac\internal\lib\transaction\tx.go -destination=D:\tugas\Github\go-rbac\internal\lib\transaction\mock\tx_mock.go
//

// Package mock_transaction is a generated GoMock package.
package mock_transaction

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
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

// Execute mocks base method.
func (m *MockManager) Execute(ctx context.Context, fn func(context.Context) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", ctx, fn)
	ret0, _ := ret[0].(error)
	return ret0
}

// Execute indicates an expected call of Execute.
func (mr *MockManagerMockRecorder) Execute(ctx, fn any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockManager)(nil).Execute), ctx, fn)
}