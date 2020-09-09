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
func New(cr domain.CartRepository, tc time.Duration) domain.CartUsecase {
	return &CartUsecase{
		repo:           cr,
		timeoutContext: tc,
	}
}
