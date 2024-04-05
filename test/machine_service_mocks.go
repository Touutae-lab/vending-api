// Code generated by MockGen. DO NOT EDIT.
// Source: .\internal\service\machine_service.go
//
// Generated by this command:
//
//	mockgen -source .\internal\service\machine_service.go -destination test/machine_service_mocks.go -package mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	uuid "github.com/google/uuid"
	model "github.com/touutae-lab/vending-api/internal/database/vending_machine/vending_machine_service/model"
	gomock "go.uber.org/mock/gomock"
	context "golang.org/x/net/context"
)

// MockMachineService is a mock of MachineService interface.
type MockMachineService struct {
	ctrl     *gomock.Controller
	recorder *MockMachineServiceMockRecorder
}

// MockMachineServiceMockRecorder is the mock recorder for MockMachineService.
type MockMachineServiceMockRecorder struct {
	mock *MockMachineService
}

// NewMockMachineService creates a new mock instance.
func NewMockMachineService(ctrl *gomock.Controller) *MockMachineService {
	mock := &MockMachineService{ctrl: ctrl}
	mock.recorder = &MockMachineServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMachineService) EXPECT() *MockMachineServiceMockRecorder {
	return m.recorder
}

// CreateMachine mocks base method.
func (m *MockMachineService) CreateMachine(ctx context.Context, machine model.Machine) (uuid.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateMachine", ctx, machine)
	ret0, _ := ret[0].(uuid.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateMachine indicates an expected call of CreateMachine.
func (mr *MockMachineServiceMockRecorder) CreateMachine(ctx, machine any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMachine", reflect.TypeOf((*MockMachineService)(nil).CreateMachine), ctx, machine)
}

// DeleteMachine mocks base method.
func (m *MockMachineService) DeleteMachine(ctx context.Context, id uuid.UUID) (uuid.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteMachine", ctx, id)
	ret0, _ := ret[0].(uuid.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteMachine indicates an expected call of DeleteMachine.
func (mr *MockMachineServiceMockRecorder) DeleteMachine(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMachine", reflect.TypeOf((*MockMachineService)(nil).DeleteMachine), ctx, id)
}

// GetAllMachine mocks base method.
func (m *MockMachineService) GetAllMachine(ctx context.Context) ([]model.Machine, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllMachine", ctx)
	ret0, _ := ret[0].([]model.Machine)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllMachine indicates an expected call of GetAllMachine.
func (mr *MockMachineServiceMockRecorder) GetAllMachine(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllMachine", reflect.TypeOf((*MockMachineService)(nil).GetAllMachine), ctx)
}

// GetMachineByID mocks base method.
func (m *MockMachineService) GetMachineByID(ctx context.Context, id uuid.UUID) (model.Machine, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMachineByID", ctx, id)
	ret0, _ := ret[0].(model.Machine)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMachineByID indicates an expected call of GetMachineByID.
func (mr *MockMachineServiceMockRecorder) GetMachineByID(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMachineByID", reflect.TypeOf((*MockMachineService)(nil).GetMachineByID), ctx, id)
}

// GetMachineTypeByID mocks base method.
func (m *MockMachineService) GetMachineTypeByID(ctx context.Context, id int32) (model.Machinetype, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMachineTypeByID", ctx, id)
	ret0, _ := ret[0].(model.Machinetype)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMachineTypeByID indicates an expected call of GetMachineTypeByID.
func (mr *MockMachineServiceMockRecorder) GetMachineTypeByID(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMachineTypeByID", reflect.TypeOf((*MockMachineService)(nil).GetMachineTypeByID), ctx, id)
}

// UpdateMachine mocks base method.
func (m *MockMachineService) UpdateMachine(ctx context.Context, machine model.Machine) (uuid.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateMachine", ctx, machine)
	ret0, _ := ret[0].(uuid.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateMachine indicates an expected call of UpdateMachine.
func (mr *MockMachineServiceMockRecorder) UpdateMachine(ctx, machine any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMachine", reflect.TypeOf((*MockMachineService)(nil).UpdateMachine), ctx, machine)
}
