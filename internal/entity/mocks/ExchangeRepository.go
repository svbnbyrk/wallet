// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	entity "github.com/svbnbyrk/wallet/internal/entity"
)

// ExchangeRepository is an autogenerated mock type for the ExchangeRepository type
type ExchangeRepository struct {
	mock.Mock
}

// GetByCurrency provides a mock function with given fields: ctx, currency
func (_m *ExchangeRepository) GetByCurrency(ctx context.Context, currency string) (entity.Exchange, error) {
	ret := _m.Called(ctx, currency)

	var r0 entity.Exchange
	if rf, ok := ret.Get(0).(func(context.Context, string) entity.Exchange); ok {
		r0 = rf(ctx, currency)
	} else {
		r0 = ret.Get(0).(entity.Exchange)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, currency)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewExchangeRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewExchangeRepository creates a new instance of ExchangeRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewExchangeRepository(t mockConstructorTestingTNewExchangeRepository) *ExchangeRepository {
	mock := &ExchangeRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
