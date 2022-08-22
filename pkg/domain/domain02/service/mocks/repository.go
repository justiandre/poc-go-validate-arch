// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	domain02 "github.com/justiandre/poc-go-validate-arch/pkg/domain/domain02"
	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the repository type
type Repository struct {
	mock.Mock
}

// Save provides a mock function with given fields: entity
func (_m *Repository) Save(entity domain02.Domain02) string {
	ret := _m.Called(entity)

	var r0 string
	if rf, ok := ret.Get(0).(func(domain02.Domain02) string); ok {
		r0 = rf(entity)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

type mockConstructorTestingTNewRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewRepository creates a new instance of Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRepository(t mockConstructorTestingTNewRepository) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}