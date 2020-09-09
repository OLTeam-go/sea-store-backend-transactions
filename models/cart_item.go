package models

import (
	"time"

	"github.com/google/uuid"
)

//CartItem represent the cart_items model in the database
type CartItem struct {
	Base
	ItemID    uuid.UUID  `json:"item_id" gorm:"type:uuid"`
	CartID    uuid.UUID  `json:"cart_id" gorm:"type:uuid"`
	Quantity  int        `json:"quantity"`
	DeletedAt *time.Time `json:"deleted_at"`
}
