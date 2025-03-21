package dtos

import "github.com/ppcamp/go-pismo-code-challenge/pkg/enums"

type CreateTransaction struct {
	AccountId       int64               `json:"account_id"`
	OperationTypeId enums.OperationType `json:"operation_type_id"`
	Amount          float64             `json:"amount"`
}
