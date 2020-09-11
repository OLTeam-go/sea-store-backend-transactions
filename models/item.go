package models

import (
	"time"

	"github.com/google/uuid"
)

// Item represent the item model from api
type Item struct {
	ID          uuid.UUID  `json:"id"`
	MerchantID  uuid.UUID  `json:"merchant_id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Category    string     `json:"category"`
	Price       float64    `json:"price"`
	Quantity    int        `json:"quantity"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}
