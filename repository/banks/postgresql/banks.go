package postgresql

import (
	"context"

	"github.com/OLTeam-go/sea-store-backend-transactions/models"
)

func (br *bankPostgresqlRepository) Fetch(ctx context.Context) (*[]models.Bank, error) {
	return nil, nil
}
