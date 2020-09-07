package usecase

import (
	"time"

	"github.com/OLTeam-go/sea-store-backend-transactions/domain"
)

type bankUsecase struct {
	repo           domain.BankRepository
	timeoutContext time.Duration
}

//New function initializze usecase for services
func New(br domain.BankRepository, tc time.Duration) domain.BankUsecase {
	return &bankUsecase{
		repo:           br,
		timeoutContext: tc,
	}
}
