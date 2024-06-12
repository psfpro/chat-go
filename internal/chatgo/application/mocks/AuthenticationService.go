// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	uuid "github.com/gofrs/uuid"
	mock "github.com/stretchr/testify/mock"
)

// AuthenticationService is an autogenerated mock type for the AuthenticationService type
type AuthenticationService struct {
	mock.Mock
}

type AuthenticationService_Expecter struct {
	mock *mock.Mock
}

func (_m *AuthenticationService) EXPECT() *AuthenticationService_Expecter {
	return &AuthenticationService_Expecter{mock: &_m.Mock}
}

// AccessToken provides a mock function with given fields: userID
func (_m *AuthenticationService) AccessToken(userID uuid.UUID) (string, error) {
	ret := _m.Called(userID)

	if len(ret) == 0 {
		panic("no return value specified for AccessToken")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(uuid.UUID) (string, error)); ok {
		return rf(userID)
	}
	if rf, ok := ret.Get(0).(func(uuid.UUID) string); ok {
		r0 = rf(userID)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(uuid.UUID) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AuthenticationService_AccessToken_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AccessToken'
type AuthenticationService_AccessToken_Call struct {
	*mock.Call
}

// AccessToken is a helper method to define mock.On call
//   - userID uuid.UUID
func (_e *AuthenticationService_Expecter) AccessToken(userID interface{}) *AuthenticationService_AccessToken_Call {
	return &AuthenticationService_AccessToken_Call{Call: _e.mock.On("AccessToken", userID)}
}

func (_c *AuthenticationService_AccessToken_Call) Run(run func(userID uuid.UUID)) *AuthenticationService_AccessToken_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uuid.UUID))
	})
	return _c
}

func (_c *AuthenticationService_AccessToken_Call) Return(_a0 string, _a1 error) *AuthenticationService_AccessToken_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AuthenticationService_AccessToken_Call) RunAndReturn(run func(uuid.UUID) (string, error)) *AuthenticationService_AccessToken_Call {
	_c.Call.Return(run)
	return _c
}

// CheckPassword provides a mock function with given fields: passwordHash, providedPassword
func (_m *AuthenticationService) CheckPassword(passwordHash string, providedPassword string) error {
	ret := _m.Called(passwordHash, providedPassword)

	if len(ret) == 0 {
		panic("no return value specified for CheckPassword")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(passwordHash, providedPassword)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AuthenticationService_CheckPassword_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CheckPassword'
type AuthenticationService_CheckPassword_Call struct {
	*mock.Call
}

// CheckPassword is a helper method to define mock.On call
//   - passwordHash string
//   - providedPassword string
func (_e *AuthenticationService_Expecter) CheckPassword(passwordHash interface{}, providedPassword interface{}) *AuthenticationService_CheckPassword_Call {
	return &AuthenticationService_CheckPassword_Call{Call: _e.mock.On("CheckPassword", passwordHash, providedPassword)}
}

func (_c *AuthenticationService_CheckPassword_Call) Run(run func(passwordHash string, providedPassword string)) *AuthenticationService_CheckPassword_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *AuthenticationService_CheckPassword_Call) Return(_a0 error) *AuthenticationService_CheckPassword_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AuthenticationService_CheckPassword_Call) RunAndReturn(run func(string, string) error) *AuthenticationService_CheckPassword_Call {
	_c.Call.Return(run)
	return _c
}

// GetUserID provides a mock function with given fields: tokenString
func (_m *AuthenticationService) GetUserID(tokenString string) (uuid.UUID, error) {
	ret := _m.Called(tokenString)

	if len(ret) == 0 {
		panic("no return value specified for GetUserID")
	}

	var r0 uuid.UUID
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (uuid.UUID, error)); ok {
		return rf(tokenString)
	}
	if rf, ok := ret.Get(0).(func(string) uuid.UUID); ok {
		r0 = rf(tokenString)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(uuid.UUID)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(tokenString)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AuthenticationService_GetUserID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUserID'
type AuthenticationService_GetUserID_Call struct {
	*mock.Call
}

// GetUserID is a helper method to define mock.On call
//   - tokenString string
func (_e *AuthenticationService_Expecter) GetUserID(tokenString interface{}) *AuthenticationService_GetUserID_Call {
	return &AuthenticationService_GetUserID_Call{Call: _e.mock.On("GetUserID", tokenString)}
}

func (_c *AuthenticationService_GetUserID_Call) Run(run func(tokenString string)) *AuthenticationService_GetUserID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *AuthenticationService_GetUserID_Call) Return(_a0 uuid.UUID, _a1 error) *AuthenticationService_GetUserID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AuthenticationService_GetUserID_Call) RunAndReturn(run func(string) (uuid.UUID, error)) *AuthenticationService_GetUserID_Call {
	_c.Call.Return(run)
	return _c
}

// HashPassword provides a mock function with given fields: password
func (_m *AuthenticationService) HashPassword(password string) (string, error) {
	ret := _m.Called(password)

	if len(ret) == 0 {
		panic("no return value specified for HashPassword")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(password)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(password)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AuthenticationService_HashPassword_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'HashPassword'
type AuthenticationService_HashPassword_Call struct {
	*mock.Call
}

// HashPassword is a helper method to define mock.On call
//   - password string
func (_e *AuthenticationService_Expecter) HashPassword(password interface{}) *AuthenticationService_HashPassword_Call {
	return &AuthenticationService_HashPassword_Call{Call: _e.mock.On("HashPassword", password)}
}

func (_c *AuthenticationService_HashPassword_Call) Run(run func(password string)) *AuthenticationService_HashPassword_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *AuthenticationService_HashPassword_Call) Return(_a0 string, _a1 error) *AuthenticationService_HashPassword_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AuthenticationService_HashPassword_Call) RunAndReturn(run func(string) (string, error)) *AuthenticationService_HashPassword_Call {
	_c.Call.Return(run)
	return _c
}

// RefreshToken provides a mock function with given fields:
func (_m *AuthenticationService) RefreshToken() (string, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for RefreshToken")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func() (string, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AuthenticationService_RefreshToken_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RefreshToken'
type AuthenticationService_RefreshToken_Call struct {
	*mock.Call
}

// RefreshToken is a helper method to define mock.On call
func (_e *AuthenticationService_Expecter) RefreshToken() *AuthenticationService_RefreshToken_Call {
	return &AuthenticationService_RefreshToken_Call{Call: _e.mock.On("RefreshToken")}
}

func (_c *AuthenticationService_RefreshToken_Call) Run(run func()) *AuthenticationService_RefreshToken_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *AuthenticationService_RefreshToken_Call) Return(_a0 string, _a1 error) *AuthenticationService_RefreshToken_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AuthenticationService_RefreshToken_Call) RunAndReturn(run func() (string, error)) *AuthenticationService_RefreshToken_Call {
	_c.Call.Return(run)
	return _c
}

// NewAuthenticationService creates a new instance of AuthenticationService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAuthenticationService(t interface {
	mock.TestingT
	Cleanup(func())
}) *AuthenticationService {
	mock := &AuthenticationService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
