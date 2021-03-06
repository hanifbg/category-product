// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	product "github.com/hanifbg/category-product/service/product"
	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// GetDetail provides a mock function with given fields: ProductId
func (_m *Repository) GetDetail(ProductId int) (product.Product, error) {
	ret := _m.Called(ProductId)

	var r0 product.Product
	if rf, ok := ret.Get(0).(func(int) product.Product); ok {
		r0 = rf(ProductId)
	} else {
		r0 = ret.Get(0).(product.Product)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(ProductId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProduct provides a mock function with given fields: CategoryId, param
func (_m *Repository) GetProduct(CategoryId int, param string) ([]product.Product, error) {
	ret := _m.Called(CategoryId, param)

	var r0 []product.Product
	if rf, ok := ret.Get(0).(func(int, string) []product.Product); ok {
		r0 = rf(CategoryId, param)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]product.Product)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, string) error); ok {
		r1 = rf(CategoryId, param)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
