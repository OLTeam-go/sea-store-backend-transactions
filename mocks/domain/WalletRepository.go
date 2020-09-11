// Code generated by mockery v2.2.1. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// WalletRepository is an autogenerated mock type for the WalletRepository type
type WalletRepository struct {
	mock.Mock
}

// UpdateMerchantWallet provides a mock function with given fields: c, merchantID, money
func (_m *WalletRepository) UpdateMerchantWallet(c context.Context, merchantID uuid.UUID, money float32) error {
	ret := _m.Called(c, merchantID, money)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, float32) error); ok {
		r0 = rf(c, merchantID, money)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
