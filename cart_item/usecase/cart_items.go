package usecase

import (
	"context"

	"github.com/google/uuid"
)

func (cu *cartItemUsecase) AddItemToCart(c context.Context, customerID uuid.UUID, itemID uuid.UUID) error {
	ctx, cancel := context.WithTimeout(c, cu.timeoutContext)
	defer cancel()

	cart, err := cu.cartRepo.GetActiveByCustomerID(ctx, customerID)
	if err != nil {
		return err
	}

	err = cu.cartItemRepo.AddItemToCart(ctx, cart.ID, itemID)

	return err
}

func (cu *cartItemUsecase) RemoveItemFromCart(c context.Context, customerID uuid.UUID, itemID uuid.UUID) error {
	ctx, cancel := context.WithTimeout(c, cu.timeoutContext)
	defer cancel()

	cart, err := cu.cartRepo.GetActiveByCustomerID(ctx, customerID)
	if err != nil {
		return err
	}

	err = cu.cartItemRepo.RemoveItemFromCart(ctx, cart.ID, itemID)

	return err
}
