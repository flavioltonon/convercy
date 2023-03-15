// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	valueobject "convercy/domain/valueobject"
)

// CurrencyCodeValidationService is an autogenerated mock type for the CurrencyCodeValidationService type
type CurrencyCodeValidationService struct {
	mock.Mock
}

// ValidateCurrencyCode provides a mock function with given fields: code
func (_m *CurrencyCodeValidationService) ValidateCurrencyCode(code valueobject.CurrencyCode) error {
	ret := _m.Called(code)

	var r0 error
	if rf, ok := ret.Get(0).(func(valueobject.CurrencyCode) error); ok {
		r0 = rf(code)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewCurrencyCodeValidationService interface {
	mock.TestingT
	Cleanup(func())
}

// NewCurrencyCodeValidationService creates a new instance of CurrencyCodeValidationService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCurrencyCodeValidationService(t mockConstructorTestingTNewCurrencyCodeValidationService) *CurrencyCodeValidationService {
	mock := &CurrencyCodeValidationService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
