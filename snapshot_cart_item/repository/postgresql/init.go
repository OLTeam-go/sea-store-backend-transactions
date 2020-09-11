package postgresql

import (
	"github.com/OLTeam-go/sea-store-backend-transactions/domain"
	"github.com/jinzhu/gorm"
)

type snapshotCartItemRepository struct {
	Conn     *gorm.DB
	pagesize int
}

//New function initialize repository for snapshot_cart_items
func New(Conn *gorm.DB, pagesize int) domain.SnapshotCartItemRepository {
	return &snapshotCartItemRepository{
		Conn,
		pagesize,
	}
}
