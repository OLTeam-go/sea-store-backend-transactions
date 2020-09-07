package postgresql

import (
	"github.com/OLTeam-go/sea-store-backend-transactions/repository/banks"
	"github.com/go-pg/pg"
)

type bankPostgresqlRepository struct {
	Conn *pg.DB
}

// New function instantiate object that represent bank repository
func New(Conn *pg.DB) banks.Repository {
	return &bankPostgresqlRepository{
		Conn,
	}
}
