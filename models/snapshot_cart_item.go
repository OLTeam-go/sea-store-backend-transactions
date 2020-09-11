package models

import (
	"github.com/google/uuid"
)

// SnapshotCartItem represent the snapshot_cart_items in the database
type SnapshotCartItem struct {
	Base
	ItemID        uuid.UUID `json:"item_id" gorm:"type:uuid"`
	CartID        uuid.UUID `json:"cart_id" gorm:"type:uuid"`
	TransactionID uuid.UUID `json:"transaction_id" gorm:"type:uuid"`
	MerchantID    uuid.UUID `json:"merchant_id" gorm:"type:uuid"`
	Name          string    `json:"name"`
	Category      string    `json:"category"`
	Price         float64   `json:"price"`
	Quantity      int       `json:"quantity"`
	Paid          bool      `json:"paid"`
}
