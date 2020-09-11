package usecase

import (
	"time"

	"github.com/OLTeam-go/sea-store-backend-transactions/domain"
)

//CartUsecase represent cart's usecase
type CartUsecase struct {
	repo           domain.CartRepository
	timeoutContext time.Duration
}

//New function initialize usecase used for cart
func New(r domain.AvailableRepository, tc time.Duration) domain.CartUsecase {
	return &CartUsecase{
		repo:           r.CartRepo,
		timeoutContext: tc,
	}
}
