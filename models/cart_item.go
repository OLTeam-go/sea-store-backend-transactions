package models

import (
	"time"

	"github.com/google/uuid"
)

//CartItem represent the cart_items model in the database
type CartItem struct {
	ID        uuid.UUID `json:"id" pg:"type:uuid,default:gen_random_uuid(),pk"`
	ItemID    uuid.UUID `json:"item_id" pg:"type:uuid"`
	CartID    uuid.UUID `json:"cart_id" pg:"type:uuid"`
	Quantity  int       `json:"quantity"`
	CreatedAt time.Time `json:"created_at" pg:"default:now()"`
	UpdatedAt time.Time `json:"updated_at" pg:"default:now()"`
	DeletedAt time.Time `json:"deleted_at" pg:"soft_delete"`
}
