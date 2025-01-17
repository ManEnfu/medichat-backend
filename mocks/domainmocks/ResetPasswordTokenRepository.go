// Code generated by mockery v2.10.4. DO NOT EDIT.

package domainmocks

import (
	context "context"
	domain "medichat-be/domain"

	mock "github.com/stretchr/testify/mock"
)

// ResetPasswordTokenRepository is an autogenerated mock type for the ResetPasswordTokenRepository type
type ResetPasswordTokenRepository struct {
	mock.Mock
}

// Add provides a mock function with given fields: ctx, token
func (_m *ResetPasswordTokenRepository) Add(ctx context.Context, token domain.ResetPasswordToken) (domain.ResetPasswordToken, error) {
	ret := _m.Called(ctx, token)

	var r0 domain.ResetPasswordToken
	if rf, ok := ret.Get(0).(func(context.Context, domain.ResetPasswordToken) domain.ResetPasswordToken); ok {
		r0 = rf(ctx, token)
	} else {
		r0 = ret.Get(0).(domain.ResetPasswordToken)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, domain.ResetPasswordToken) error); ok {
		r1 = rf(ctx, token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: ctx, id
func (_m *ResetPasswordTokenRepository) GetByID(ctx context.Context, id int64) (domain.ResetPasswordToken, error) {
	ret := _m.Called(ctx, id)

	var r0 domain.ResetPasswordToken
	if rf, ok := ret.Get(0).(func(context.Context, int64) domain.ResetPasswordToken); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(domain.ResetPasswordToken)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByTokenStr provides a mock function with given fields: ctx, tokenStr
func (_m *ResetPasswordTokenRepository) GetByTokenStr(ctx context.Context, tokenStr string) (domain.ResetPasswordToken, error) {
	ret := _m.Called(ctx, tokenStr)

	var r0 domain.ResetPasswordToken
	if rf, ok := ret.Get(0).(func(context.Context, string) domain.ResetPasswordToken); ok {
		r0 = rf(ctx, tokenStr)
	} else {
		r0 = ret.Get(0).(domain.ResetPasswordToken)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, tokenStr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByTokenStrAndLock provides a mock function with given fields: ctx, tokenStr
func (_m *ResetPasswordTokenRepository) GetByTokenStrAndLock(ctx context.Context, tokenStr string) (domain.ResetPasswordToken, error) {
	ret := _m.Called(ctx, tokenStr)

	var r0 domain.ResetPasswordToken
	if rf, ok := ret.Get(0).(func(context.Context, string) domain.ResetPasswordToken); ok {
		r0 = rf(ctx, tokenStr)
	} else {
		r0 = ret.Get(0).(domain.ResetPasswordToken)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, tokenStr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SoftDeleteByAccountID provides a mock function with given fields: ctx, id
func (_m *ResetPasswordTokenRepository) SoftDeleteByAccountID(ctx context.Context, id int64) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SoftDeleteByID provides a mock function with given fields: ctx, id
func (_m *ResetPasswordTokenRepository) SoftDeleteByID(ctx context.Context, id int64) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
