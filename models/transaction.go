package models

import (
	"github.com/OLTeam-go/sea-store-backend-transactions/enum"
	"github.com/google/uuid"
)

//Transaction represent the transactions model in the database
type Transaction struct {
	Base
	CustomerID        uuid.UUID              `json:"customer_id" gorm:"type:uuid"`
	CartID            uuid.UUID              `json:"cart_id" gorm:"type:uuid"`
	BankID            uuid.UUID              `json:"bank_id" gorm:"type:uuid"`
	BankAccountNumber string                 `json:"bank_account_number" gorm:"type:uuid"`
	Status            enum.TransactionStatus `json:"status"`
	Cost              float64                `json:"cost"`
	Bank              *Bank                  `json:"bank,omitempty" gorm:"foreignKey:BankID"`
	Cart              *Cart                  `json:"cart,omitempty" gorm:"foreignKey:CartID"`
	SnapshotCartItems []SnapshotCartItem     `json:"snapshot_cart_items" gorm:"foreignKey:TransactionID"`
}
