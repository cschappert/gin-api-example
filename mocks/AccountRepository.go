// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	api "github.com/cschappert/gin-api-example/pkg"
	mock "github.com/stretchr/testify/mock"
)

// AccountRepository is an autogenerated mock type for the AccountRepository type
type AccountRepository struct {
	mock.Mock
}

// CreateAccount provides a mock function with given fields: a
func (_m *AccountRepository) CreateAccount(a *api.Account) error {
	ret := _m.Called(a)

	var r0 error
	if rf, ok := ret.Get(0).(func(*api.Account) error); ok {
		r0 = rf(a)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteAccount provides a mock function with given fields: id
func (_m *AccountRepository) DeleteAccount(id int) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAccount provides a mock function with given fields: id
func (_m *AccountRepository) GetAccount(id int) (*api.Account, error) {
	ret := _m.Called(id)

	var r0 *api.Account
	if rf, ok := ret.Get(0).(func(int) *api.Account); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*api.Account)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListAccounts provides a mock function with given fields:
func (_m *AccountRepository) ListAccounts() ([]*api.Account, error) {
	ret := _m.Called()

	var r0 []*api.Account
	if rf, ok := ret.Get(0).(func() []*api.Account); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*api.Account)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewAccountRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewAccountRepository creates a new instance of AccountRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAccountRepository(t mockConstructorTestingTNewAccountRepository) *AccountRepository {
	mock := &AccountRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
