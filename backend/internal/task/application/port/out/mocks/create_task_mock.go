// Code generated by MockGen. DO NOT EDIT.
// Source: create_task.go
//
// Generated by this command:
//
//	mockgen -source=create_task.go -typed=true -destination=./mocks/create_task_mock.go
//

// Package mock_out is a generated GoMock package.
package mock_out

import (
	out "ToDo/internal/task/application/port/out"
	pkg "ToDo/pkg"
	context "context"
	reflect "reflect"

	uuid "github.com/google/uuid"
	gomock "go.uber.org/mock/gomock"
)

// MockCreateTaskPort is a mock of CreateTaskPort interface.
type MockCreateTaskPort[T pkg.Tx] struct {
	ctrl     *gomock.Controller
	recorder *MockCreateTaskPortMockRecorder[T]
	isgomock struct{}
}

// MockCreateTaskPortMockRecorder is the mock recorder for MockCreateTaskPort.
type MockCreateTaskPortMockRecorder[T pkg.Tx] struct {
	mock *MockCreateTaskPort[T]
}

// NewMockCreateTaskPort creates a new mock instance.
func NewMockCreateTaskPort[T pkg.Tx](ctrl *gomock.Controller) *MockCreateTaskPort[T] {
	mock := &MockCreateTaskPort[T]{ctrl: ctrl}
	mock.recorder = &MockCreateTaskPortMockRecorder[T]{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCreateTaskPort[T]) EXPECT() *MockCreateTaskPortMockRecorder[T] {
	return m.recorder
}

// Create mocks base method.
func (m *MockCreateTaskPort[T]) Create(ctx context.Context, dto *out.CreateTaskDTO, tx T) (uuid.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, dto, tx)
	ret0, _ := ret[0].(uuid.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockCreateTaskPortMockRecorder[T]) Create(ctx, dto, tx any) *MockCreateTaskPortCreateCall[T] {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockCreateTaskPort[T])(nil).Create), ctx, dto, tx)
	return &MockCreateTaskPortCreateCall[T]{Call: call}
}

// MockCreateTaskPortCreateCall wrap *gomock.Call
type MockCreateTaskPortCreateCall[T pkg.Tx] struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockCreateTaskPortCreateCall[T]) Return(arg0 uuid.UUID, arg1 error) *MockCreateTaskPortCreateCall[T] {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockCreateTaskPortCreateCall[T]) Do(f func(context.Context, *out.CreateTaskDTO, T) (uuid.UUID, error)) *MockCreateTaskPortCreateCall[T] {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockCreateTaskPortCreateCall[T]) DoAndReturn(f func(context.Context, *out.CreateTaskDTO, T) (uuid.UUID, error)) *MockCreateTaskPortCreateCall[T] {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
