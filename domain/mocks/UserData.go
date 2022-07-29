// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	domain "github.com/AltaProject/AltaSocialMedia/domain"
	mock "github.com/stretchr/testify/mock"
)

// UserData is an autogenerated mock type for the UserData type
type UserData struct {
	mock.Mock
}

// DeleteUser provides a mock function with given fields: userId
func (_m *UserData) DeleteUser(userId int) bool {
	ret := _m.Called(userId)

	var r0 bool
	if rf, ok := ret.Get(0).(func(int) bool); ok {
		r0 = rf(userId)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// GetSpecificUser provides a mock function with given fields: userId
func (_m *UserData) GetSpecificUser(userId int) (domain.User, error) {
	ret := _m.Called(userId)

	var r0 domain.User
	if rf, ok := ret.Get(0).(func(int) domain.User); ok {
		r0 = rf(userId)
	} else {
		r0 = ret.Get(0).(domain.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Login provides a mock function with given fields: email, password
func (_m *UserData) Login(email string, password string) (string, string, error) {
	ret := _m.Called(email, password)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, string) string); ok {
		r0 = rf(email, password)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 string
	if rf, ok := ret.Get(1).(func(string, string) string); ok {
		r1 = rf(email, password)
	} else {
		r1 = ret.Get(1).(string)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string, string) error); ok {
		r2 = rf(email, password)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Register provides a mock function with given fields: newUser
func (_m *UserData) Register(newUser domain.User) (domain.User, error) {
	ret := _m.Called(newUser)

	var r0 domain.User
	if rf, ok := ret.Get(0).(func(domain.User) domain.User); ok {
		r0 = rf(newUser)
	} else {
		r0 = ret.Get(0).(domain.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(domain.User) error); ok {
		r1 = rf(newUser)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateUser provides a mock function with given fields: updateUser, userId
func (_m *UserData) UpdateUser(updateUser domain.User, userId int) (domain.User, error) {
	ret := _m.Called(updateUser, userId)

	var r0 domain.User
	if rf, ok := ret.Get(0).(func(domain.User, int) domain.User); ok {
		r0 = rf(updateUser, userId)
	} else {
		r0 = ret.Get(0).(domain.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(domain.User, int) error); ok {
		r1 = rf(updateUser, userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewUserData interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserData creates a new instance of UserData. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserData(t mockConstructorTestingTNewUserData) *UserData {
	mock := &UserData{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}