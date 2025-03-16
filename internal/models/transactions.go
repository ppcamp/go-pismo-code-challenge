package models

type Transactions struct {
	Id          int64
	AccountId   int64
	OperationId int64
	Amount      float64
	EventDate   string
}
