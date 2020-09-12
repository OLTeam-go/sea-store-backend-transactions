package usecase

import (
	"time"

	"github.com/OLTeam-go/sea-store-backend-transactions/domain"
)

type transactionUsecase struct {
	itemRepo        domain.ItemRepository
	transactionRepo domain.TransactionRepository
	cartRepo        domain.CartRepository
	snapshotRepo    domain.SnapshotCartItemRepository
	walletRepo      domain.WalletRepository
	timeoutContext  time.Duration
}

//New function initialize the usecase for transactions
func New(r domain.AvailableRepository, tc time.Duration) domain.TransactionUsecase {
	return &transactionUsecase{
		itemRepo:        r.ItemRepo,
		cartRepo:        r.CartRepo,
		transactionRepo: r.TransactionRepo,
		snapshotRepo:    r.SnapshotRepo,
		walletRepo:      r.WalletRepo,
		timeoutContext:  tc,
	}
}
