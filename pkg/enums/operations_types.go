package enums

// OperationType is just a shortcut for the operation_type_table_ref
type OperationType int64

const (
	OpNormalPurchase OperationType = iota + 1
	OpPurchaseInstallments
	OpWithdrawl
	OpCreditVoucher
)
