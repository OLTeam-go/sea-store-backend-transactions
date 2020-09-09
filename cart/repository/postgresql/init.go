package postgresql

import (
	"github.com/OLTeam-go/sea-store-backend-transactions/domain"
	"github.com/jinzhu/gorm"
)

type cartRepository struct {
	Conn     *gorm.DB
	pagesize int
}

// New initialize cart repository
func New(Conn *gorm.DB, pagesize int) domain.CartRepository {
	return &cartRepository{
		Conn,
		pagesize,
	}
}
