package domain

import (
	"context"

	"github.com/OLTeam-go/sea-store-backend-transactions/models"
	"github.com/google/uuid"
)

//ItemRepository represent item's repository
type ItemRepository interface {
	FetchByIDs(c context.Context, id []uuid.UUID) ([]*models.Item, error)
	UpdateQuantity(c context.Context, id []uuid.UUID, action string) error
}
