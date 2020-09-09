package postgresql

import (
	"context"
	"time"

	"github.com/OLTeam-go/sea-store-backend-transactions/models"
	"github.com/google/uuid"
)

func (cr *cartItemRepository) AddItemToCart(ctx context.Context, cartID uuid.UUID, itemID uuid.UUID) error {
	cartItem := models.CartItem{
		CartID:   cartID,
		ItemID:   itemID,
		Quantity: 1,
	}
	var cartInstance models.CartItem
	DB := cr.Conn.FirstOrCreate(&cartInstance, cartItem)

	return DB.Error
}

func (cr *cartItemRepository) RemoveItemFromCart(ctx context.Context, cartID uuid.UUID, itemID uuid.UUID) error {
	cartItem := models.CartItem{
		CartID:   cartID,
		ItemID:   itemID,
		Quantity: 1,
	}
	var cartInstance models.CartItem
	cr.Conn.LogMode(true)
	DB := cr.Conn.First(&cartInstance, cartItem)

	if DB.RecordNotFound() == false {
		now := time.Now()
		cartInstance.DeletedAt = &now
		DB = cr.Conn.Save(&cartInstance)
		return DB.Error
	}

	return DB.Error
}
