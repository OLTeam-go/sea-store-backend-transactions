package models

import (
	"time"

	"github.com/google/uuid"
)

// Cart represent the carts model in the database
type Cart struct {
	ID         uuid.UUID `json:"id" pg:"type:uuid,default:gen_random_uuid(),pk"`
	CustomerID uuid.UUID `json:"customer_id" pg:"type:uuid"`
	Active     bool      `json:"active"`
	CreatedAt  time.Time `json:"created_at" pg:"default:now()"`
	UpdatedAt  time.Time `json:"updated_at" pg:"default:now()"`
}
