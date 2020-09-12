package usecase

import (
	"context"
	"errors"

	"github.com/OLTeam-go/sea-store-backend-transactions/models"
	"github.com/google/uuid"
)

func (u *snapshotCartItemUsecase) FetchSnapshotCartItemsByMerchantID(c context.Context, page int, merchantID uuid.UUID) ([]*models.SnapshotCartItem, error) {
	if page < 0 {
		return nil, errors.New("page is invalid")
	}
	ctx, cancel := context.WithTimeout(c, u.timeoutContext)
	defer cancel()

	res, err := u.snapshotRepo.FetchPaidItemsByMerchantID(ctx, page, merchantID)
	if err != nil {
		return nil, err
	}

	return res, nil
}
