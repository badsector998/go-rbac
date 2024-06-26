// Code generated by MockGen. DO NOT EDIT.
// Source: ./burjo_group.go
//
// Generated by this command:
//
//	mockgen -source ./burjo_group.go -destination ./mock/burjo_group_mock.go
//

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	context "context"
	reflect "reflect"

	domain "github.com/badsector998/go-rbac/internal/domain"
	gomock "go.uber.org/mock/gomock"
)

// MockBurjoGroupRepository is a mock of BurjoGroupRepository interface.
type MockBurjoGroupRepository struct {
	ctrl     *gomock.Controller
	recorder *MockBurjoGroupRepositoryMockRecorder
}

// MockBurjoGroupRepositoryMockRecorder is the mock recorder for MockBurjoGroupRepository.
type MockBurjoGroupRepositoryMockRecorder struct {
	mock *MockBurjoGroupRepository
}

// NewMockBurjoGroupRepository creates a new mock instance.
func NewMockBurjoGroupRepository(ctrl *gomock.Controller) *MockBurjoGroupRepository {
	mock := &MockBurjoGroupRepository{ctrl: ctrl}
	mock.recorder = &MockBurjoGroupRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBurjoGroupRepository) EXPECT() *MockBurjoGroupRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockBurjoGroupRepository) Create(ctx context.Context, bg domain.BurjoGroup) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, bg)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockBurjoGroupRepositoryMockRecorder) Create(ctx, bg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockBurjoGroupRepository)(nil).Create), ctx, bg)
}

// Delete mocks base method.
func (m *MockBurjoGroupRepository) Delete(ctx context.Context, bg domain.BurjoGroup) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, bg)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockBurjoGroupRepositoryMockRecorder) Delete(ctx, bg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockBurjoGroupRepository)(nil).Delete), ctx, bg)
}

// ReadAll mocks base method.
func (m *MockBurjoGroupRepository) ReadAll(ctx context.Context) ([]domain.BurjoGroup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadAll", ctx)
	ret0, _ := ret[0].([]domain.BurjoGroup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadAll indicates an expected call of ReadAll.
func (mr *MockBurjoGroupRepositoryMockRecorder) ReadAll(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadAll", reflect.TypeOf((*MockBurjoGroupRepository)(nil).ReadAll), ctx)
}

// ReadByID mocks base method.
func (m *MockBurjoGroupRepository) ReadByID(ctx context.Context, id uint) (domain.BurjoGroup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadByID", ctx, id)
	ret0, _ := ret[0].(domain.BurjoGroup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadByID indicates an expected call of ReadByID.
func (mr *MockBurjoGroupRepositoryMockRecorder) ReadByID(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadByID", reflect.TypeOf((*MockBurjoGroupRepository)(nil).ReadByID), ctx, id)
}

// Update mocks base method.
func (m *MockBurjoGroupRepository) Update(ctx context.Context, bg domain.BurjoGroup) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, bg)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockBurjoGroupRepositoryMockRecorder) Update(ctx, bg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockBurjoGroupRepository)(nil).Update), ctx, bg)
}
