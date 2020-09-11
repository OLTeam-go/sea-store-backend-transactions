package postgresql

import (
	"context"

	"github.com/OLTeam-go/sea-store-backend-transactions/models"
)

func (r *snapshotCartItemRepository) Create(c context.Context, s models.SnapshotCartItem) (*models.SnapshotCartItem, error) {
	DB := r.Conn.Create(&s)
	if DB.Error != nil {
		return nil, DB.Error
	}
	return &s, nil
}
