// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	domain01 "github.com/justiandre/poc-go-validate-arch/pkg/domain/domain01"
	mock "github.com/stretchr/testify/mock"
)

// Service is an autogenerated mock type for the service type
type Service struct {
	mock.Mock
}

// DeleteByID provides a mock function with given fields: id
func (_m *Service) DeleteByID(id string) {
	_m.Called(id)
}

// FindByID provides a mock function with given fields: id
func (_m *Service) FindByID(id string) domain01.Domain01 {
	ret := _m.Called(id)

	var r0 domain01.Domain01
	if rf, ok := ret.Get(0).(func(string) domain01.Domain01); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(domain01.Domain01)
	}

	return r0
}

// Save provides a mock function with given fields: entity
func (_m *Service) Save(entity domain01.Domain01) string {
	ret := _m.Called(entity)

	var r0 string
	if rf, ok := ret.Get(0).(func(domain01.Domain01) string); ok {
		r0 = rf(entity)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

type mockConstructorTestingTNewService interface {
	mock.TestingT
	Cleanup(func())
}

// NewService creates a new instance of Service. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewService(t mockConstructorTestingTNewService) *Service {
	mock := &Service{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
