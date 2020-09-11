package postgresql

import (
	"github.com/OLTeam-go/sea-store-backend-transactions/domain"
	"github.com/jinzhu/gorm"
)

type transactionRepository struct {
	Conn     *gorm.DB
	pagesize int
}

//New function initialize the repository for transaction
func New(Conn *gorm.DB, pagesize int) domain.TransactionRepository {
	return &transactionRepository{
		Conn,
		pagesize,
	}
}
