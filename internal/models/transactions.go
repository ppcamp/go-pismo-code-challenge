package models

import "github.com/ppcamp/go-pismo-code-challenge/pkg/enums"

type Transaction struct {
	Id          int64
	AccountId   int64
	OperationId enums.OperationType
	Amount      float64
	EventDate   string
}
