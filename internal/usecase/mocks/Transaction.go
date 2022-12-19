// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	entity "github.com/svbnbyrk/wallet/internal/entity"
)

// Transaction is an autogenerated mock type for the Transaction type
type Transaction struct {
	mock.Mock
}

// History provides a mock function with given fields: ctx
func (_m *Transaction) History(ctx context.Context) ([]entity.Transaction, error) {
	ret := _m.Called(ctx)

	var r0 []entity.Transaction
	if rf, ok := ret.Get(0).(func(context.Context) []entity.Transaction); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Post provides a mock function with given fields: ctx, t
func (_m *Transaction) Post(ctx context.Context, t entity.Transaction) error {
	ret := _m.Called(ctx, t)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, entity.Transaction) error); ok {
		r0 = rf(ctx, t)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewTransaction interface {
	mock.TestingT
	Cleanup(func())
}

// NewTransaction creates a new instance of Transaction. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTransaction(t mockConstructorTestingTNewTransaction) *Transaction {
	mock := &Transaction{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
