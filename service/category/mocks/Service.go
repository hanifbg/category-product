// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	category "github.com/hanifbg/category-product/service/category"
	mock "github.com/stretchr/testify/mock"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// GetCategory provides a mock function with given fields:
func (_m *Service) GetCategory() ([]category.Category, error) {
	ret := _m.Called()

	var r0 []category.Category
	if rf, ok := ret.Get(0).(func() []category.Category); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]category.Category)
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
