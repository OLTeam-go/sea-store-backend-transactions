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
func New(r domain.AvailableRepository, tc time.Duration) domain.BankUsecase {
	return &bankUsecase{
		repo:           r.BankRepo,
		timeoutContext: tc,
	}
}
