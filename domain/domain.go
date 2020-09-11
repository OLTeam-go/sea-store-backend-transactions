package domain

//AvailableRepository are container of all available repository
type AvailableRepository struct {
	BankRepo        BankRepository
	CartRepo        CartRepository
	CartItemRepo    CartItemRepository
	TransactionRepo TransactionRepository
	SnapshotRepo    SnapshotCartItemRepository
	ItemRepo        ItemRepository
}

//AvailableUsecase are container of all available usecase
type AvailableUsecase struct {
	BankUsecase             BankUsecase
	CartUsecase             CartUsecase
	TransactionUsecase      TransactionUsecase
	CartItemUsecase         CartItemUsecase
	SnapshotCartItemUsecase SnapshotCartItemUsecase
}
