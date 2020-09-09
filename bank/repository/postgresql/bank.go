package postgresql

import (
	"context"

	"github.com/OLTeam-go/sea-store-backend-transactions/models"
)

func (br *bankPostgresqlRepository) Fetch(ctx context.Context) ([]*models.Bank, error) {
	var bank []*models.Bank
	err := br.Conn.Find(&bank, "active = true")
	if err.Error != nil {
		return nil, err.Error
	}

	return bank, nil
}
