package models

import (
	"time"

	"github.com/google/uuid"
)

//Transaction represent the transactions model in the database
type Transaction struct {
	ID                uuid.UUID `json:"id" pg:"type:uuid,default:gen_random_uuid(),pk"`
	CustomerID        uuid.UUID `json:"customer_id" pg:"type:uuid"`
	CartID            uuid.UUID `json:"cart_id" pg:"type:uuid"`
	BankID            uuid.UUID `json:"bank_id" pg:"type:uuid"`
	BankAccountNumber string    `json:"bank_account_number" pg:"type:uuid"`
	Status            bool      `json:"status"`
	Cost              float64   `json:"cost"`
	CreatedAt         time.Time `json:"created_at" pg:"default:now()"`
	UpdatedAt         time.Time `json:"updated_at" pg:"default:now()"`
}
