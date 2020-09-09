package postgresql

import (
	"github.com/OLTeam-go/sea-store-backend-transactions/domain"
	"github.com/jinzhu/gorm"
)

type bankPostgresqlRepository struct {
	Conn *gorm.DB
}

// New function instantiate object that represent bank repository
func New(Conn *gorm.DB) domain.BankRepository {
	return &bankPostgresqlRepository{
		Conn,
	}
}
