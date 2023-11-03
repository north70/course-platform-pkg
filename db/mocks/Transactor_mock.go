// Code generated by mockery v2.36.0. DO NOT EDIT.

package db

import (
	context "context"

	pgx "github.com/jackc/pgx/v4"
	mock "github.com/stretchr/testify/mock"
)

// MockTransactor is an autogenerated mock type for the Transactor type
type MockTransactor struct {
	mock.Mock
}

type MockTransactor_Expecter struct {
	mock *mock.Mock
}

func (_m *MockTransactor) EXPECT() *MockTransactor_Expecter {
	return &MockTransactor_Expecter{mock: &_m.Mock}
}

// BeginTx provides a mock function with given fields: ctx, txOptions
func (_m *MockTransactor) BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error) {
	ret := _m.Called(ctx, txOptions)

	var r0 pgx.Tx
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, pgx.TxOptions) (pgx.Tx, error)); ok {
		return rf(ctx, txOptions)
	}
	if rf, ok := ret.Get(0).(func(context.Context, pgx.TxOptions) pgx.Tx); ok {
		r0 = rf(ctx, txOptions)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(pgx.Tx)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, pgx.TxOptions) error); ok {
		r1 = rf(ctx, txOptions)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTransactor_BeginTx_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BeginTx'
type MockTransactor_BeginTx_Call struct {
	*mock.Call
}

// BeginTx is a helper method to define mock.On call
//   - ctx context.Context
//   - txOptions pgx.TxOptions
func (_e *MockTransactor_Expecter) BeginTx(ctx interface{}, txOptions interface{}) *MockTransactor_BeginTx_Call {
	return &MockTransactor_BeginTx_Call{Call: _e.mock.On("BeginTx", ctx, txOptions)}
}

func (_c *MockTransactor_BeginTx_Call) Run(run func(ctx context.Context, txOptions pgx.TxOptions)) *MockTransactor_BeginTx_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(pgx.TxOptions))
	})
	return _c
}

func (_c *MockTransactor_BeginTx_Call) Return(_a0 pgx.Tx, _a1 error) *MockTransactor_BeginTx_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockTransactor_BeginTx_Call) RunAndReturn(run func(context.Context, pgx.TxOptions) (pgx.Tx, error)) *MockTransactor_BeginTx_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockTransactor creates a new instance of MockTransactor. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockTransactor(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockTransactor {
	mock := &MockTransactor{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}