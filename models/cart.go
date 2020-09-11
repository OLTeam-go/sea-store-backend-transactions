package models

import (
	"github.com/google/uuid"
)

// Cart represent the carts model in the database
type Cart struct {
	Base
	CustomerID        uuid.UUID          `json:"customer_id" gorm:"type:uuid"`
	Active            bool               `json:"active" gorm:"type:boolean;default:true"`
	CartItems         []CartItem         `json:"cart_items,omitempty" gorm:"foreignKey:CartID"`
	SnapshotCartItems []SnapshotCartItem `json:"snapshot_cart_items,omitempty" gorm:"foreignKey:CartID"`
}
