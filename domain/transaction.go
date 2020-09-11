package domain

import (
	"context"

	"github.com/OLTeam-go/sea-store-backend-transactions/enum"

	"github.com/OLTeam-go/sea-store-backend-transactions/models"
	"github.com/google/uuid"
)

//TransactionRepository represent the contract for transaction
type TransactionRepository interface {
	CreateTransaction(c context.Context, t models.Transaction) (*models.Transaction, error)
	UpdateStatusTransaction(c context.Context, id uuid.UUID, s enum.TransactionStatus) error
	GetByID(c context.Context, id uuid.UUID) (*models.Transaction, error)
	FetchTransactions(c context.Context, page int, status []enum.TransactionStatus) ([]*models.Transaction, error)
	FetchTransactionsByCustomerID(c context.Context, page int, customerID uuid.UUID, status []enum.TransactionStatus) ([]*models.Transaction, error)
}

//TransactionUsecase represent the business logic for transaction
type TransactionUsecase interface {
	CreateTransaction(c context.Context, customerID uuid.UUID, bankID uuid.UUID, bankAccountNumber string) (*models.Transaction, error)
	FetchTransactions(c context.Context, page int, status enum.TransactionFilterStatus) ([]*models.Transaction, error)
	FetchTransactionsByCustomerID(c context.Context, page int, customerID uuid.UUID, status enum.TransactionFilterStatus) ([]*models.Transaction, error)
	AcceptStatusTransaction(c context.Context, id uuid.UUID) error
	RejectStatusTransaction(c context.Context, id uuid.UUID) error
}
