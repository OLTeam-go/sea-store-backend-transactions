// Code generated by mockery v2.2.1. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// CartItemRepository is an autogenerated mock type for the CartItemRepository type
type CartItemRepository struct {
	mock.Mock
}

// AddItemToCart provides a mock function with given fields: ctx, cartID, itemID
func (_m *CartItemRepository) AddItemToCart(ctx context.Context, cartID uuid.UUID, itemID uuid.UUID) error {
	ret := _m.Called(ctx, cartID, itemID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, uuid.UUID) error); ok {
		r0 = rf(ctx, cartID, itemID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RemoveItemFromCart provides a mock function with given fields: ctx, cartID, itemID
func (_m *CartItemRepository) RemoveItemFromCart(ctx context.Context, cartID uuid.UUID, itemID uuid.UUID) error {
	ret := _m.Called(ctx, cartID, itemID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, uuid.UUID) error); ok {
		r0 = rf(ctx, cartID, itemID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
