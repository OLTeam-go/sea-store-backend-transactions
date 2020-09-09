package usecase

import (
	"context"

	"github.com/OLTeam-go/sea-store-backend-transactions/models"
	"github.com/google/uuid"
)

//GetActiveByCustomerID return user's active cart based on user's ID
func (cu *CartUsecase) GetActiveByCustomerID(c context.Context, id uuid.UUID) (*models.Cart, error) {
	ctx, cancel := context.WithTimeout(c, cu.timeoutContext)
	defer cancel()

	res, err := cu.repo.GetActiveByCustomerID(ctx, id)

	if err != nil {
		return nil, err
	}

	return res, nil
}

//FetchHistoryByCustomerID fetch paginated customer's inactive cart
func (cu *CartUsecase) FetchHistoryByCustomerID(c context.Context, id uuid.UUID, page int) ([]*models.Cart, error) {
	ctx, cancel := context.WithTimeout(c, cu.timeoutContext)
	defer cancel()

	res, err := cu.repo.FetchHistoryByCustomerID(ctx, id, page)

	if err != nil {
		return nil, err
	}

	return res, nil
}
