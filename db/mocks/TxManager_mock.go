// Code generated by mockery v2.36.0. DO NOT EDIT.

package db

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockTxManager is an autogenerated mock type for the TxManager type
type MockTxManager struct {
	mock.Mock
}

type MockTxManager_Expecter struct {
	mock *mock.Mock
}

func (_m *MockTxManager) EXPECT() *MockTxManager_Expecter {
	return &MockTxManager_Expecter{mock: &_m.Mock}
}

// ReadCommitted provides a mock function with given fields: ctx, f
func (_m *MockTxManager) ReadCommitted(ctx context.Context, f func(context.Context) error) error {
	ret := _m.Called(ctx, f)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, func(context.Context) error) error); ok {
		r0 = rf(ctx, f)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockTxManager_ReadCommitted_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ReadCommitted'
type MockTxManager_ReadCommitted_Call struct {
	*mock.Call
}

// ReadCommitted is a helper method to define mock.On call
//   - ctx context.Context
//   - f func(context.Context) error
func (_e *MockTxManager_Expecter) ReadCommitted(ctx interface{}, f interface{}) *MockTxManager_ReadCommitted_Call {
	return &MockTxManager_ReadCommitted_Call{Call: _e.mock.On("ReadCommitted", ctx, f)}
}

func (_c *MockTxManager_ReadCommitted_Call) Run(run func(ctx context.Context, f func(context.Context) error)) *MockTxManager_ReadCommitted_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(func(context.Context) error))
	})
	return _c
}

func (_c *MockTxManager_ReadCommitted_Call) Return(_a0 error) *MockTxManager_ReadCommitted_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTxManager_ReadCommitted_Call) RunAndReturn(run func(context.Context, func(context.Context) error) error) *MockTxManager_ReadCommitted_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockTxManager creates a new instance of MockTxManager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockTxManager(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockTxManager {
	mock := &MockTxManager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
