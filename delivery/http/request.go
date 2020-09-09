package http

import "github.com/google/uuid"

//CartItemRequest represent the incoming request for cart_items
type CartItemRequest struct {
	ItemID   uuid.UUID `json:"item_id"`
	Quantity int       `json:"quantity"`
}
