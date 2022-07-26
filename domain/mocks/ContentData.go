// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	domain "github.com/AltaProject/AltaSocialMedia/domain"
	mock "github.com/stretchr/testify/mock"
)

// ContentData is an autogenerated mock type for the ContentData type
type ContentData struct {
	mock.Mock
}

// AddNewContent provides a mock function with given fields: newContent
func (_m *ContentData) AddNewContent(newContent domain.Content) (domain.Content, error) {
	ret := _m.Called(newContent)

	var r0 domain.Content
	if rf, ok := ret.Get(0).(func(domain.Content) domain.Content); ok {
		r0 = rf(newContent)
	} else {
		r0 = ret.Get(0).(domain.Content)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(domain.Content) error); ok {
		r1 = rf(newContent)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: contentId
func (_m *ContentData) Delete(contentId int) bool {
	ret := _m.Called(contentId)

	var r0 bool
	if rf, ok := ret.Get(0).(func(int) bool); ok {
		r0 = rf(contentId)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// GetAllContent provides a mock function with given fields:
func (_m *ContentData) GetAllContent() ([]domain.Content, error) {
	ret := _m.Called()

	var r0 []domain.Content
	if rf, ok := ret.Get(0).(func() []domain.Content); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Content)
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

// GetContentId provides a mock function with given fields: contentId
func (_m *ContentData) GetContentId(contentId int) (domain.Content, error) {
	ret := _m.Called(contentId)

	var r0 domain.Content
	if rf, ok := ret.Get(0).(func(int) domain.Content); ok {
		r0 = rf(contentId)
	} else {
		r0 = ret.Get(0).(domain.Content)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(contentId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: userId, newContent
func (_m *ContentData) Update(userId int, newContent domain.Content) (domain.Content, error) {
	ret := _m.Called(userId, newContent)

	var r0 domain.Content
	if rf, ok := ret.Get(0).(func(int, domain.Content) domain.Content); ok {
		r0 = rf(userId, newContent)
	} else {
		r0 = ret.Get(0).(domain.Content)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, domain.Content) error); ok {
		r1 = rf(userId, newContent)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewContentData interface {
	mock.TestingT
	Cleanup(func())
}

// NewContentData creates a new instance of ContentData. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewContentData(t mockConstructorTestingTNewContentData) *ContentData {
	mock := &ContentData{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
