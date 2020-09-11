package domain

import (
	"context"

	"github.com/google/uuid"
)

//WalletRepository represent wallet's repository
type WalletRepository interface {
	UpdateMerchantWallet(c context.Context, merchantID uuid.UUID, money float32) error
}
