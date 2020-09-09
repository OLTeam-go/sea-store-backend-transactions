package domain

import (
	"context"

	"github.com/OLTeam-go/sea-store-backend-transactions/models"
)

// BankRepository represent the banks repository contract
type BankRepository interface {
	Fetch(ctx context.Context) ([]*models.Bank, error)
}

// BankUsecase represent the bank usecase
type BankUsecase interface {
	Fetch(ctx context.Context) ([]*models.Bank, error)
}
