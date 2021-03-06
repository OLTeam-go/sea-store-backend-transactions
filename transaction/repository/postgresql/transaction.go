package postgresql

import (
	"context"
	"log"

	"github.com/OLTeam-go/sea-store-backend-transactions/enum"
	"github.com/jinzhu/gorm"

	"github.com/OLTeam-go/sea-store-backend-transactions/models"
	"github.com/google/uuid"
)

func (r *transactionRepository) CreateTransaction(c context.Context, t models.Transaction) (*models.Transaction, error) {
	r.Conn.LogMode(true)
	log.Println(t.SnapshotCartItems[0].CartID)
	DB := r.Conn.Create(&t)
	if DB.Error != nil {
		return nil, DB.Error
	}

	return &t, nil
}

func (r *transactionRepository) UpdateStatusTransaction(c context.Context, id uuid.UUID, s enum.TransactionStatus) error {
	DB := r.Conn.
		Model(&models.Transaction{}).
		Omit("Bank").
		Omit("Cart").
		Omit("SnapshotCartItems").
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"status": s,
		})

	return DB.Error
}

func (r *transactionRepository) GetByID(c context.Context, id uuid.UUID) (*models.Transaction, error) {
	var transaction models.Transaction
	transaction.ID = id
	DB := r.Conn.Preload("Bank").
		Preload("SnapshotCartItems").
		Preload("Cart").
		First(&transaction)
	if DB.Error != nil {
		return nil, DB.Error
	}
	return &transaction, nil
}

func (r *transactionRepository) FetchTransactions(c context.Context, page int, status []enum.TransactionStatus) ([]*models.Transaction, error) {
	var transactions []*models.Transaction
	offset := (page - 1) * r.pagesize
	limit := r.pagesize
	var st []int
	for _, s := range status {
		st = append(st, int(s))
	}

	var DB *gorm.DB
	if page != 0 {
		DB = r.Conn.
			Preload("Bank").
			Preload("SnapshotCartItems").
			Preload("Cart").
			Offset(offset).
			Limit(limit).
			Where("status IN (?)", st).
			Find(&transactions)
	} else {
		DB = r.Conn.
			Preload("Bank").
			Preload("SnapshotCartItems").
			Preload("Cart").
			Where("status IN (?)", st).
			Find(&transactions)
	}

	if DB.Error != nil {
		return nil, DB.Error
	}
	return transactions, nil
}

func (r *transactionRepository) FetchTransactionsByCustomerID(c context.Context, page int, customerID uuid.UUID, status []enum.TransactionStatus) ([]*models.Transaction, error) {
	var transactions []*models.Transaction
	offset := (page - 1) * r.pagesize
	limit := r.pagesize
	var st []int
	for _, s := range status {
		st = append(st, int(s))
	}
	var DB *gorm.DB
	if page != 0 {
		DB = r.Conn.
			Preload("Bank").
			Preload("SnapshotCartItems").
			Preload("Cart").
			Offset(offset).
			Limit(limit).
			Where("customer_id = ?", customerID).
			Where("status IN (?)", st).
			Find(&transactions)
	} else {
		DB = r.Conn.
			Preload("Bank").
			Preload("SnapshotCartItems").
			Preload("Cart").
			Where("customer_id = ?", customerID).
			Where("status IN (?)", st).
			Find(&transactions)
	}

	if DB.Error != nil {
		return nil, DB.Error
	}
	return transactions, nil
}
