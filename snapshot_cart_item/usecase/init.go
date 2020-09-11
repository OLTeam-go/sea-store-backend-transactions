package usecase

import (
	"time"

	"github.com/OLTeam-go/sea-store-backend-transactions/domain"
)

type snapshotCartItemUsecase struct {
	snapshotRepo   domain.SnapshotCartItemRepository
	timeoutContext time.Duration
}

//New function initialize repository for snapshot cart items
func New(r domain.AvailableRepository, timeoutContext time.Duration) domain.SnapshotCartItemUsecase {
	return &snapshotCartItemUsecase{
		snapshotRepo:   r.SnapshotRepo,
		timeoutContext: timeoutContext,
	}
}
