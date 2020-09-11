package domain

import (
	"context"

	"github.com/OLTeam-go/sea-store-backend-transactions/models"
	"github.com/google/uuid"
)

//SnapshotCartItemRepository represent snapshot_cart_item's repository
type SnapshotCartItemRepository interface {
	Create(c context.Context, s models.SnapshotCartItem) (*models.SnapshotCartItem, error)
	FetchPaidItemsByMerchantID(c context.Context, page int, merchantID uuid.UUID) ([]*models.SnapshotCartItem, error)
}

//SnapshotCartItemUsecase represent snapshot_cart_items's usecase
type SnapshotCartItemUsecase interface {
	FetchSnapshotCartItemsByMerchantID(c context.Context, page int, merchantID uuid.UUID) ([]*models.SnapshotCartItem, error)
}
