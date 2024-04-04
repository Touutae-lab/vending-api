// Code generated by MockGen. DO NOT EDIT.
// Source: .\internal\handler\machine_handler_interface.go
//
// Generated by this command:
//
//	mockgen -source .\internal\handler\machine_handler_interface.go -destination test/machine_handler_mocks.go -package mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	model "github.com/touutae-lab/vending-api/internal/database/vending_machine/vending_machine_service/model"
	gomock "go.uber.org/mock/gomock"
	context "golang.org/x/net/context"
)

// MockMachineHandlerInterface is a mock of MachineHandlerInterface interface.
type MockMachineHandlerInterface struct {
	ctrl     *gomock.Controller
	recorder *MockMachineHandlerInterfaceMockRecorder
}

// MockMachineHandlerInterfaceMockRecorder is the mock recorder for MockMachineHandlerInterface.
type MockMachineHandlerInterfaceMockRecorder struct {
	mock *MockMachineHandlerInterface
}

// NewMockMachineHandlerInterface creates a new mock instance.
func NewMockMachineHandlerInterface(ctrl *gomock.Controller) *MockMachineHandlerInterface {
	mock := &MockMachineHandlerInterface{ctrl: ctrl}
	mock.recorder = &MockMachineHandlerInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMachineHandlerInterface) EXPECT() *MockMachineHandlerInterfaceMockRecorder {
	return m.recorder
}

// CreateMachine mocks base method.
func (m *MockMachineHandlerInterface) CreateMachine(ctx context.Context, machine model.Machine) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateMachine", ctx, machine)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateMachine indicates an expected call of CreateMachine.
func (mr *MockMachineHandlerInterfaceMockRecorder) CreateMachine(ctx, machine any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMachine", reflect.TypeOf((*MockMachineHandlerInterface)(nil).CreateMachine), ctx, machine)
}

// DeleteMachine mocks base method.
func (m *MockMachineHandlerInterface) DeleteMachine(ctx context.Context, id int32) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteMachine", ctx, id)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteMachine indicates an expected call of DeleteMachine.
func (mr *MockMachineHandlerInterfaceMockRecorder) DeleteMachine(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMachine", reflect.TypeOf((*MockMachineHandlerInterface)(nil).DeleteMachine), ctx, id)
}

// GetAllMachine mocks base method.
func (m *MockMachineHandlerInterface) GetAllMachine(ctx context.Context) ([]model.Machine, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllMachine", ctx)
	ret0, _ := ret[0].([]model.Machine)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllMachine indicates an expected call of GetAllMachine.
func (mr *MockMachineHandlerInterfaceMockRecorder) GetAllMachine(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllMachine", reflect.TypeOf((*MockMachineHandlerInterface)(nil).GetAllMachine), ctx)
}

// GetMachineByID mocks base method.
func (m *MockMachineHandlerInterface) GetMachineByID(ctx context.Context, id int32) (model.Machine, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMachineByID", ctx, id)
	ret0, _ := ret[0].(model.Machine)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMachineByID indicates an expected call of GetMachineByID.
func (mr *MockMachineHandlerInterfaceMockRecorder) GetMachineByID(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMachineByID", reflect.TypeOf((*MockMachineHandlerInterface)(nil).GetMachineByID), ctx, id)
}

// GetMachineTypeByID mocks base method.
func (m *MockMachineHandlerInterface) GetMachineTypeByID(ctx context.Context, id int32) (model.Machinetype, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMachineTypeByID", ctx, id)
	ret0, _ := ret[0].(model.Machinetype)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMachineTypeByID indicates an expected call of GetMachineTypeByID.
func (mr *MockMachineHandlerInterfaceMockRecorder) GetMachineTypeByID(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMachineTypeByID", reflect.TypeOf((*MockMachineHandlerInterface)(nil).GetMachineTypeByID), ctx, id)
}

// UpdateMachine mocks base method.
func (m *MockMachineHandlerInterface) UpdateMachine(ctx context.Context, machine model.Machine) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateMachine", ctx, machine)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateMachine indicates an expected call of UpdateMachine.
func (mr *MockMachineHandlerInterfaceMockRecorder) UpdateMachine(ctx, machine any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMachine", reflect.TypeOf((*MockMachineHandlerInterface)(nil).UpdateMachine), ctx, machine)
}
