// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	uuid "github.com/gofrs/uuid"
	mock "github.com/stretchr/testify/mock"
)

// ChatGoService is an autogenerated mock type for the ChatGoService type
type ChatGoService struct {
	mock.Mock
}

type ChatGoService_Expecter struct {
	mock *mock.Mock
}

func (_m *ChatGoService) EXPECT() *ChatGoService_Expecter {
	return &ChatGoService_Expecter{mock: &_m.Mock}
}

// AddTask provides a mock function with given fields: title, description
func (_m *ChatGoService) AddTask(title string, description string) (uuid.UUID, error) {
	ret := _m.Called(title, description)

	if len(ret) == 0 {
		panic("no return value specified for AddTask")
	}

	var r0 uuid.UUID
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (uuid.UUID, error)); ok {
		return rf(title, description)
	}
	if rf, ok := ret.Get(0).(func(string, string) uuid.UUID); ok {
		r0 = rf(title, description)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(uuid.UUID)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(title, description)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ChatGoService_AddTask_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddTask'
type ChatGoService_AddTask_Call struct {
	*mock.Call
}

// AddTask is a helper method to define mock.On call
//   - title string
//   - description string
func (_e *ChatGoService_Expecter) AddTask(title interface{}, description interface{}) *ChatGoService_AddTask_Call {
	return &ChatGoService_AddTask_Call{Call: _e.mock.On("AddTask", title, description)}
}

func (_c *ChatGoService_AddTask_Call) Run(run func(title string, description string)) *ChatGoService_AddTask_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *ChatGoService_AddTask_Call) Return(_a0 uuid.UUID, _a1 error) *ChatGoService_AddTask_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ChatGoService_AddTask_Call) RunAndReturn(run func(string, string) (uuid.UUID, error)) *ChatGoService_AddTask_Call {
	_c.Call.Return(run)
	return _c
}

// NewChatGoService creates a new instance of ChatGoService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewChatGoService(t interface {
	mock.TestingT
	Cleanup(func())
}) *ChatGoService {
	mock := &ChatGoService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
