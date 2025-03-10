// Code generated by MockGen. DO NOT EDIT.
// Source: create_task.go
//
// Generated by this command:
//
//	mockgen -source=create_task.go -typed=true -destination=./mocks/create_task_mock.go
//

// Package mock_in is a generated GoMock package.
package mock_in

import (
	in "ToDo/internal/task/application/port/in"
	context "context"
	reflect "reflect"

	uuid "github.com/google/uuid"
	gomock "go.uber.org/mock/gomock"
)

// MockCreateTaskUseCase is a mock of CreateTaskUseCase interface.
type MockCreateTaskUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockCreateTaskUseCaseMockRecorder
	isgomock struct{}
}

// MockCreateTaskUseCaseMockRecorder is the mock recorder for MockCreateTaskUseCase.
type MockCreateTaskUseCaseMockRecorder struct {
	mock *MockCreateTaskUseCase
}

// NewMockCreateTaskUseCase creates a new mock instance.
func NewMockCreateTaskUseCase(ctrl *gomock.Controller) *MockCreateTaskUseCase {
	mock := &MockCreateTaskUseCase{ctrl: ctrl}
	mock.recorder = &MockCreateTaskUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCreateTaskUseCase) EXPECT() *MockCreateTaskUseCaseMockRecorder {
	return m.recorder
}

// Execute mocks base method.
func (m *MockCreateTaskUseCase) Execute(ctx context.Context, dto *in.CreateTaskDTO) (uuid.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", ctx, dto)
	ret0, _ := ret[0].(uuid.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Execute indicates an expected call of Execute.
func (mr *MockCreateTaskUseCaseMockRecorder) Execute(ctx, dto any) *MockCreateTaskUseCaseExecuteCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockCreateTaskUseCase)(nil).Execute), ctx, dto)
	return &MockCreateTaskUseCaseExecuteCall{Call: call}
}

// MockCreateTaskUseCaseExecuteCall wrap *gomock.Call
type MockCreateTaskUseCaseExecuteCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockCreateTaskUseCaseExecuteCall) Return(arg0 uuid.UUID, arg1 error) *MockCreateTaskUseCaseExecuteCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockCreateTaskUseCaseExecuteCall) Do(f func(context.Context, *in.CreateTaskDTO) (uuid.UUID, error)) *MockCreateTaskUseCaseExecuteCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockCreateTaskUseCaseExecuteCall) DoAndReturn(f func(context.Context, *in.CreateTaskDTO) (uuid.UUID, error)) *MockCreateTaskUseCaseExecuteCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
