package usecase

import (
	"context"

	"github.com/OLTeam-go/sea-store-backend-transactions/models"
)

func (bu *bankUsecase) Fetch(c context.Context) ([]*models.Bank, error) {
	ctx, cancel := context.WithTimeout(c, bu.timeoutContext)
	defer cancel()

	res, err := bu.repo.Fetch(ctx)

	if err != nil {
		return nil, err
	}

	return res, nil
}
