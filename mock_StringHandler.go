// Code generated by mockery v2.33.1. DO NOT EDIT.

package klib

import mock "github.com/stretchr/testify/mock"

// MockStringHandler is an autogenerated mock type for the StringHandler type
type MockStringHandler struct {
	mock.Mock
}

type MockStringHandler_Expecter struct {
	mock *mock.Mock
}

func (_m *MockStringHandler) EXPECT() *MockStringHandler_Expecter {
	return &MockStringHandler_Expecter{mock: &_m.Mock}
}

// Handle provides a mock function with given fields: input
func (_m *MockStringHandler) Handle(input string) (string, error) {
	ret := _m.Called(input)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(input)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(input)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockStringHandler_Handle_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Handle'
type MockStringHandler_Handle_Call struct {
	*mock.Call
}

// Handle is a helper method to define mock.On call
//   - input string
func (_e *MockStringHandler_Expecter) Handle(input interface{}) *MockStringHandler_Handle_Call {
	return &MockStringHandler_Handle_Call{Call: _e.mock.On("Handle", input)}
}

func (_c *MockStringHandler_Handle_Call) Run(run func(input string)) *MockStringHandler_Handle_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockStringHandler_Handle_Call) Return(_a0 string, _a1 error) *MockStringHandler_Handle_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockStringHandler_Handle_Call) RunAndReturn(run func(string) (string, error)) *MockStringHandler_Handle_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockStringHandler creates a new instance of MockStringHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockStringHandler(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockStringHandler {
	mock := &MockStringHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
