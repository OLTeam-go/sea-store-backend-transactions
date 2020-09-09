package domain

import (
	"context"

	"github.com/google/uuid"
)

// CartItemRepository represent the cart_item's repository contract
type CartItemRepository interface {
	AddItemToCart(ctx context.Context, cartID uuid.UUID, itemID uuid.UUID) error
	RemoveItemFromCart(ctx context.Context, cartID uuid.UUID, itemID uuid.UUID) error
}

// CartItemUsecase represent the cart's usecase
type CartItemUsecase interface {
	AddItemToCart(ctx context.Context, customerID uuid.UUID, itemID uuid.UUID) error
	RemoveItemFromCart(ctx context.Context, customerID uuid.UUID, itemID uuid.UUID) error
}
