// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	contracts "github.com/yfpanne/app-functions-sdk-go/v2/pkg/interfaces"

	mock "github.com/stretchr/testify/mock"
)

// StoreClient is an autogenerated mock type for the StoreClient type
type StoreClient struct {
	mock.Mock
}

// Disconnect provides a mock function with given fields:
func (_m *StoreClient) Disconnect() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RemoveFromStore provides a mock function with given fields: o
func (_m *StoreClient) RemoveFromStore(o contracts.StoredObject) error {
	ret := _m.Called(o)

	var r0 error
	if rf, ok := ret.Get(0).(func(contracts.StoredObject) error); ok {
		r0 = rf(o)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RetrieveFromStore provides a mock function with given fields: appServiceKey
func (_m *StoreClient) RetrieveFromStore(appServiceKey string) ([]contracts.StoredObject, error) {
	ret := _m.Called(appServiceKey)

	var r0 []contracts.StoredObject
	if rf, ok := ret.Get(0).(func(string) []contracts.StoredObject); ok {
		r0 = rf(appServiceKey)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]contracts.StoredObject)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(appServiceKey)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Store provides a mock function with given fields: o
func (_m *StoreClient) Store(o contracts.StoredObject) (string, error) {
	ret := _m.Called(o)

	var r0 string
	var r1 error

	if rf, ok := ret.Get(0).(func(contracts.StoredObject) (string, error)); ok {
		r0, r1 = rf(o)
	} else {
		r1 = ret.Error(0)
	}

	return r0, r1
}

// Update provides a mock function with given fields: o
func (_m *StoreClient) Update(o contracts.StoredObject) error {
	ret := _m.Called(o)

	var r0 error
	if rf, ok := ret.Get(0).(func(contracts.StoredObject) error); ok {
		r0 = rf(o)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
