package models

import (
	"github.com/google/uuid"
)

// Cart represent the carts model in the database
type Cart struct {
	Base
	CustomerID uuid.UUID  `json:"customer_id" gorm:"type:uuid"`
	Active     bool       `json:"active"`
	CartItems  []CartItem `json:"cart_items" gorm:"foreignKey:CartID"`
}
