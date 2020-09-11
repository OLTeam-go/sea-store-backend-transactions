package postgresql

import (
	"context"
	"time"

	"github.com/OLTeam-go/sea-store-backend-transactions/models"
	"github.com/google/uuid"
)

func (r *snapshotCartItemRepository) Create(c context.Context, s models.SnapshotCartItem) (*models.SnapshotCartItem, error) {
	DB := r.Conn.Create(&s)
	if DB.Error != nil {
		return nil, DB.Error
	}
	return &s, nil
}

func (r *snapshotCartItemRepository) FetchPaidItemsByMerchantID(c context.Context, page int, merchantID uuid.UUID) ([]*models.SnapshotCartItem, error) {
	var snapshotCartItems []*models.SnapshotCartItem
	offset := (page - 1) * r.pagesize
	limit := r.pagesize
	r.Conn.LogMode(true)
	DB := r.Conn.
		Where("merchant_id = ?", merchantID).
		Where("paid = true").
		Offset(offset).
		Limit(limit).
		Find(&snapshotCartItems)

	if DB.Error != nil {
		return nil, DB.Error
	}

	return snapshotCartItems, nil
}

func (r *snapshotCartItemRepository) SetPaid(c context.Context, s models.SnapshotCartItem) error {
	s.UpdatedAt = time.Now()
	s.Paid = true
	DB := r.Conn.Save(&s)
	if DB.Error != nil {
		return DB.Error
	}
	return nil
}
