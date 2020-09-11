package postgresql

import (
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/OLTeam-go/sea-store-backend-transactions/models"
)

func (cr *cartRepository) calculateOffsetLimit(page int) (int, int, error) {
	if page < 0 {
		return 0, 0, models.ErrBadParamInput
	}

	return (page - 1) * cr.pagesize, cr.pagesize, nil
}

// GetActiveByCustomerID return active cart for Customer based on its ID
func (cr *cartRepository) GetActiveByCustomerID(ctx context.Context, id uuid.UUID) (*models.Cart, error) {
	var cart models.Cart

	DB := cr.Conn.Preload("CartItems").FirstOrCreate(&cart, models.Cart{
		CustomerID: id,
		Active:     true,
	})
	if DB.Error != nil {
		return nil, DB.Error
	}
	return &cart, nil
}

// FetchHistoryByCustomerID return paginated history cart for customer based on its ID
func (cr *cartRepository) FetchHistoryByCustomerID(ctx context.Context, id uuid.UUID, page int) ([]*models.Cart, error) {
	offset, limit, err := cr.calculateOffsetLimit(page)
	if err != nil {
		return nil, err
	}

	var cart []*models.Cart

	DB := cr.Conn.Preload("SnapshotCartItems").Find(&cart, "customer_id = ? and active = false", id).Offset(offset).Limit(limit)

	if DB.Error != nil {
		return nil, DB.Error
	}

	return cart, nil
}

func (cr *cartRepository) Update(ctx context.Context, c models.Cart) error {
	c.UpdatedAt = time.Now()
	var cart models.Cart
	DB := cr.Conn.Omit("SnapshotCartItems").Omit("CartItems").Model(&cart).Updates(map[string]interface{}{
		"Active":     c.Active,
		"UpdatedAt":  c.UpdatedAt,
		"CustomerID": c.CustomerID,
	})
	if DB.Error != nil {
		return DB.Error
	}
	return nil
}

func (cr *cartRepository) GetByID(ctx context.Context, id uuid.UUID) (*models.Cart, error) {
	var cart models.Cart
	DB := cr.Conn.
		Preload("CartItems").
		Preload("SnapshotCartItems").
		First(&cart, "id = ?", id)
	if DB.Error != nil {
		return nil, DB.Error
	}

	return &cart, nil
}
