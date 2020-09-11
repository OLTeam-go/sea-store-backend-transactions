package enum

//TransactionStatus represent the status of a transaction
type TransactionStatus int

//TransactionFilterStatus represent all available filter for transaction
type TransactionFilterStatus string

const (
	//TransactionRejected represent the accepted rejected
	TransactionRejected TransactionStatus = iota
	//TransactionPending represent the accepted pending
	TransactionPending
	//TransactionAccepted represent the accepted status
	TransactionAccepted
)

const (
	//TransactionFilterRejected represent rejected only transactions
	TransactionFilterRejected TransactionFilterStatus = "rejected"
	//TransactionFilterAll represent all transactions
	TransactionFilterAll TransactionFilterStatus = "all"
	//TransactionFilterPending represent pending only transactions
	TransactionFilterPending TransactionFilterStatus = "pending"
	//TransactionFilterAccepted represent accepted only transactions
	TransactionFilterAccepted TransactionFilterStatus = "accepted"
)
