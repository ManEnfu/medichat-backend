// Code generated by mockery v2.10.4. DO NOT EDIT.

package domainmocks

import (
	context "context"
	domain "medichat-be/domain"

	mock "github.com/stretchr/testify/mock"
)

// VerifyEmailTokenRepository is an autogenerated mock type for the VerifyEmailTokenRepository type
type VerifyEmailTokenRepository struct {
	mock.Mock
}

// Add provides a mock function with given fields: ctx, token
func (_m *VerifyEmailTokenRepository) Add(ctx context.Context, token domain.VerifyEmailToken) (domain.VerifyEmailToken, error) {
	ret := _m.Called(ctx, token)

	var r0 domain.VerifyEmailToken
	if rf, ok := ret.Get(0).(func(context.Context, domain.VerifyEmailToken) domain.VerifyEmailToken); ok {
		r0 = rf(ctx, token)
	} else {
		r0 = ret.Get(0).(domain.VerifyEmailToken)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, domain.VerifyEmailToken) error); ok {
		r1 = rf(ctx, token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: ctx, id
func (_m *VerifyEmailTokenRepository) GetByID(ctx context.Context, id int64) (domain.VerifyEmailToken, error) {
	ret := _m.Called(ctx, id)

	var r0 domain.VerifyEmailToken
	if rf, ok := ret.Get(0).(func(context.Context, int64) domain.VerifyEmailToken); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(domain.VerifyEmailToken)
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
func (_m *VerifyEmailTokenRepository) GetByTokenStr(ctx context.Context, tokenStr string) (domain.VerifyEmailToken, error) {
	ret := _m.Called(ctx, tokenStr)

	var r0 domain.VerifyEmailToken
	if rf, ok := ret.Get(0).(func(context.Context, string) domain.VerifyEmailToken); ok {
		r0 = rf(ctx, tokenStr)
	} else {
		r0 = ret.Get(0).(domain.VerifyEmailToken)
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
func (_m *VerifyEmailTokenRepository) GetByTokenStrAndLock(ctx context.Context, tokenStr string) (domain.VerifyEmailToken, error) {
	ret := _m.Called(ctx, tokenStr)

	var r0 domain.VerifyEmailToken
	if rf, ok := ret.Get(0).(func(context.Context, string) domain.VerifyEmailToken); ok {
		r0 = rf(ctx, tokenStr)
	} else {
		r0 = ret.Get(0).(domain.VerifyEmailToken)
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
func (_m *VerifyEmailTokenRepository) SoftDeleteByAccountID(ctx context.Context, id int64) error {
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
func (_m *VerifyEmailTokenRepository) SoftDeleteByID(ctx context.Context, id int64) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
