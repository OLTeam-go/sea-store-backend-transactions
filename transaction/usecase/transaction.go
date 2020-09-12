package usecase

import (
	"context"
	"errors"

	"github.com/OLTeam-go/sea-store-backend-transactions/enum"

	"github.com/OLTeam-go/sea-store-backend-transactions/models"
	"github.com/google/uuid"
)

func getItemsIDFromCart(cart *models.Cart) []uuid.UUID {
	var ids []uuid.UUID
	for _, cartItem := range cart.CartItems {
		ids = append(ids, cartItem.ItemID)
	}
	return ids
}

func createCostAndSnapshotItems(cart *models.Cart, items []*models.Item) (float64, []models.SnapshotCartItem) {
	var cost float64
	var snapshotCartItems []models.SnapshotCartItem
	cost = 0
	for _, item := range items {
		cost = cost + item.Price
		snapshotItem := models.SnapshotCartItem{
			ItemID:     item.ID,
			CartID:     cart.ID,
			MerchantID: item.MerchantID,
			Name:       item.Name,
			Category:   item.Category,
			Price:      item.Price,
			Quantity:   1,
		}
		snapshotCartItems = append(snapshotCartItems, snapshotItem)
	}
	return cost, snapshotCartItems
}

func (u *transactionUsecase) CreateTransaction(c context.Context, customerID uuid.UUID, bankID uuid.UUID, bankAccountNumber string) (*models.Transaction, error) {
	ctx, cancel := context.WithTimeout(c, u.timeoutContext)
	defer cancel()

	cart, err := u.cartRepo.GetActiveByCustomerID(ctx, customerID)
	if err != nil {
		return nil, err
	}
	if cart.CartItems == nil {
		return nil, errors.New("cart is empty")
	}

	ids := getItemsIDFromCart(cart)
	items, err := u.itemRepo.FetchByIDs(ctx, ids)
	if err != nil {
		return nil, err
	}

	cost, snapshotCartItems := createCostAndSnapshotItems(cart, items)
	transaction := models.Transaction{
		BankID:            bankID,
		BankAccountNumber: bankAccountNumber,
		CartID:            cart.ID,
		Status:            enum.TransactionPending,
		Cost:              cost,
		CustomerID:        customerID,
		SnapshotCartItems: snapshotCartItems,
	}
	res, err := u.transactionRepo.CreateTransaction(ctx, transaction)
	if err != nil {
		return nil, err
	}

	cart.Active = false
	err = u.cartRepo.Update(ctx, *cart)
	if err != nil {
		return nil, err
	}
	err = u.itemRepo.UpdateQuantity(ctx, ids, "sold")
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (u *transactionUsecase) AcceptStatusTransaction(c context.Context, id uuid.UUID) error {
	ctx, cancel := context.WithTimeout(c, u.timeoutContext)
	defer cancel()

	transaction, err := u.transactionRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	// if transaction.Status != enum.TransactionPending {
	// 	return errors.New("Transaction is not a pending transaction")
	// }

	for _, snapshotItems := range transaction.SnapshotCartItems {
		_ = u.walletRepo.UpdateMerchantWallet(ctx, snapshotItems.MerchantID, snapshotItems.Price)
		_ = u.snapshotRepo.SetPaid(ctx, snapshotItems)
	}

	err = u.transactionRepo.UpdateStatusTransaction(ctx, id, enum.TransactionAccepted)

	return err
}

func (u *transactionUsecase) RejectStatusTransaction(c context.Context, id uuid.UUID) error {
	ctx, cancel := context.WithTimeout(c, u.timeoutContext)
	defer cancel()

	transaction, err := u.transactionRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	// if transaction.Status != enum.TransactionPending {
	// 	return errors.New("Transaction is not a pending transaction")
	// }
	cartID := transaction.CartID
	cart, err := u.cartRepo.GetByID(ctx, cartID)
	if err != nil {
		return err
	}

	ids := getItemsIDFromCart(cart)

	err = u.itemRepo.UpdateQuantity(ctx, ids, "available")
	if err != nil {
		return err
	}

	err = u.transactionRepo.UpdateStatusTransaction(ctx, id, enum.TransactionRejected)

	return err
}

func getTransactionStatusFromFilter(status enum.TransactionFilterStatus) []enum.TransactionStatus {
	switch status {
	case enum.TransactionFilterAll:
		return []enum.TransactionStatus{enum.TransactionAccepted, enum.TransactionRejected, enum.TransactionPending}
	case enum.TransactionFilterAccepted:
		return []enum.TransactionStatus{enum.TransactionAccepted}
	case enum.TransactionFilterRejected:
		return []enum.TransactionStatus{enum.TransactionRejected}
	case enum.TransactionFilterPending:
		return []enum.TransactionStatus{enum.TransactionPending}
	default:
		return []enum.TransactionStatus{enum.TransactionAccepted, enum.TransactionRejected, enum.TransactionPending}
	}
}

func (u *transactionUsecase) FetchTransactions(c context.Context, page int, filter enum.TransactionFilterStatus) ([]*models.Transaction, error) {
	ctx, cancel := context.WithTimeout(c, u.timeoutContext)
	defer cancel()
	status := getTransactionStatusFromFilter(filter)

	if page < 0 {
		return nil, errors.New("page is invalid")
	}

	res, err := u.transactionRepo.FetchTransactions(ctx, page, status)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (u *transactionUsecase) FetchTransactionsByCustomerID(c context.Context, page int, customerID uuid.UUID, filter enum.TransactionFilterStatus) ([]*models.Transaction, error) {
	ctx, cancel := context.WithTimeout(c, u.timeoutContext)
	defer cancel()
	status := getTransactionStatusFromFilter(filter)

	if page < 0 {
		return nil, errors.New("page is invalid")
	}

	res, err := u.transactionRepo.FetchTransactionsByCustomerID(ctx, page, customerID, status)

	if err != nil {
		return nil, err
	}

	return res, nil
}
