package api

import (
	"github.com/OLTeam-go/sea-store-backend-transactions/domain"
)

type walletRepository struct {
	APIURL string
}

//New function initialize repository for wallet
func New(APIURL string) domain.WalletRepository {
	return &walletRepository{
		APIURL,
	}
}
