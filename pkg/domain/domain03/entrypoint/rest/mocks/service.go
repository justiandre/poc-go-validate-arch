// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	domain02 "github.com/justiandre/poc-go-validate-arch/pkg/domain/domain02"
	domain03 "github.com/justiandre/poc-go-validate-arch/pkg/domain/domain03"

	mock "github.com/stretchr/testify/mock"
)

// Service is an autogenerated mock type for the service type
type Service struct {
	mock.Mock
}

// Save provides a mock function with given fields: entity, entityDomain02
func (_m *Service) Save(entity domain03.Domain03, entityDomain02 domain02.Domain02) string {
	ret := _m.Called(entity, entityDomain02)

	var r0 string
	if rf, ok := ret.Get(0).(func(domain03.Domain03, domain02.Domain02) string); ok {
		r0 = rf(entity, entityDomain02)
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
