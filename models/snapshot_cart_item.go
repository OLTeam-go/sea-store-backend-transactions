package models

import (
	"time"

	"github.com/google/uuid"
)

// SnapshotCartItem represent the snapshot_cart_items in the database
type SnapshotCartItem struct {
	ID         uuid.UUID `json:"id" pg:"type:uuid,default:gen_random_uuid(),pk"`
	ItemID     uuid.UUID `json:"item_id" pg:"type:uuid"`
	CartID     uuid.UUID `json:"cart_id" pg:"type:uuid"`
	MerchantID uuid.UUID `json:"merchant_id" pg:"type:uuid"`
	Name       string    `json:"name"`
	Category   string    `json:"category"`
	Price      string    `json:"price"`
	Quantity   int       `json:"quantity"`
	CreatedAt  time.Time `json:"created_at" pg:"default:now()"`
	UpdatedAt  time.Time `json:"updated_at" pg:"default:now()"`
}
