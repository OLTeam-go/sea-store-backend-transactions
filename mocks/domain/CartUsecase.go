// Code generated by mockery v2.2.1. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	models "github.com/OLTeam-go/sea-store-backend-transactions/models"

	uuid "github.com/google/uuid"
)

// CartUsecase is an autogenerated mock type for the CartUsecase type
type CartUsecase struct {
	mock.Mock
}

// FetchHistoryByCustomerID provides a mock function with given fields: ctx, id, page
func (_m *CartUsecase) FetchHistoryByCustomerID(ctx context.Context, id uuid.UUID, page int) ([]*models.Cart, error) {
	ret := _m.Called(ctx, id, page)

	var r0 []*models.Cart
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, int) []*models.Cart); ok {
		r0 = rf(ctx, id, page)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Cart)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID, int) error); ok {
		r1 = rf(ctx, id, page)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetActiveByCustomerID provides a mock function with given fields: ctx, id
func (_m *CartUsecase) GetActiveByCustomerID(ctx context.Context, id uuid.UUID) (*models.Cart, error) {
	ret := _m.Called(ctx, id)

	var r0 *models.Cart
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) *models.Cart); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Cart)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
