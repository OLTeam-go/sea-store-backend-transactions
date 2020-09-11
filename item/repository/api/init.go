package api

import (
	"github.com/OLTeam-go/sea-store-backend-transactions/domain"
)

type itemRepository struct {
	APIURL string
}

//New function initialize the repository for items
func New(URL string) domain.ItemRepository {
	return &itemRepository{
		APIURL: URL,
	}
}
