package banks

import (
	"context"

	"github.com/OLTeam-go/sea-store-backend-transactions/models"
)

// Repository represent the banks repository contract
type Repository interface {
	Fetch(ctx context.Context) (*[]models.Bank, error)
}
