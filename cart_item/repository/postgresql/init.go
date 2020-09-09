package postgresql

import (
	"github.com/OLTeam-go/sea-store-backend-transactions/domain"
	"github.com/jinzhu/gorm"
)

type cartItemRepository struct {
	Conn     *gorm.DB
	pagesize int
}

//New function initialize cart_item's repository
func New(Conn *gorm.DB, pagesize int) domain.CartItemRepository {
	return &cartItemRepository{
		Conn,
		pagesize,
	}
}
