package constants

const (
	SuccessMessage      = "Success"
	ErrFailedBadRequest = "Data tidak sesuai"
	ErrServerError      = "Terjadi kesalahan pada server"
)

const (
	TransactionStatusPending  = "PENDING"
	TransactionStatusSuccess  = "SUCCESS"
	TransactionStatusFailed   = "FAILED"
	TransactionStatusReversed = "REVERSED"
)

const (
	TransactionTypeTopUp    = "TOPUP"
	TransactionTypePurchase = "PURCHASE"
	TransactionTypeRefund   = "REFUND"
)

var MapTransactionType = map[string]bool{
	TransactionTypeTopUp:    true,
	TransactionTypePurchase: true,
	TransactionTypeRefund:   true,
}
