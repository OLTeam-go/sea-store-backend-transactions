package postgresql

import (
	"context"

	"github.com/OLTeam-go/sea-store-backend-transactions/models"
)

func (br *bankPostgresqlRepository) Fetch(ctx context.Context) (*[]models.Bank, error) {
	var bank []models.Bank
	err := br.Conn.Model(&bank).Where("active = true").Returning("id", "name").Select()
	if err != nil {
		return nil, err
	}

	return &bank, nil
}
