// Code generated by mockery v2.2.1. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	models "github.com/OLTeam-go/sea-store-backend-transactions/models"
)

// SnapshotCartItemRepository is an autogenerated mock type for the SnapshotCartItemRepository type
type SnapshotCartItemRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: c, s
func (_m *SnapshotCartItemRepository) Create(c context.Context, s models.SnapshotCartItem) (*models.SnapshotCartItem, error) {
	ret := _m.Called(c, s)

	var r0 *models.SnapshotCartItem
	if rf, ok := ret.Get(0).(func(context.Context, models.SnapshotCartItem) *models.SnapshotCartItem); ok {
		r0 = rf(c, s)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.SnapshotCartItem)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, models.SnapshotCartItem) error); ok {
		r1 = rf(c, s)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
