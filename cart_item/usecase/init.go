package usecase

import (
	"time"

	"github.com/OLTeam-go/sea-store-backend-transactions/domain"
)

type cartItemUsecase struct {
	cartItemRepo   domain.CartItemRepository
	cartRepo       domain.CartRepository
	timeoutContext time.Duration
}

//New function initialize cart_item's usecase
func New(r domain.AvailableRepository, timeoutContext time.Duration) domain.CartItemUsecase {
	return &cartItemUsecase{
		cartItemRepo:   r.CartItemRepo,
		cartRepo:       r.CartRepo,
		timeoutContext: timeoutContext,
	}
}
