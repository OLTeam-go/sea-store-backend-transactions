package domain

import (
	"context"

	"github.com/OLTeam-go/sea-store-backend-transactions/models"
	"github.com/google/uuid"
)

// CartRepository represent the cart's repository contract
type CartRepository interface {
	GetActiveByCustomerID(ctx context.Context, id uuid.UUID) (*models.Cart, error)
	FetchHistoryByCustomerID(ctx context.Context, id uuid.UUID, page int) ([]*models.Cart, error)
}

// CartUsecase represent the cart's usecase
type CartUsecase interface {
	GetActiveByCustomerID(ctx context.Context, id uuid.UUID) (*models.Cart, error)
	FetchHistoryByCustomerID(ctx context.Context, id uuid.UUID, page int) ([]*models.Cart, error)
}
