package domain

import (
	"context"

	"github.com/OLTeam-go/sea-store-backend-transactions/models"
)

//SnapshotCartItemRepository represent snapshot_cart_item's repository
type SnapshotCartItemRepository interface {
	Create(c context.Context, s models.SnapshotCartItem) (*models.SnapshotCartItem, error)
}
