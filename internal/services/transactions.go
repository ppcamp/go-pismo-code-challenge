package handlers

import (
	"context"

	"github.com/ppcamp/go-pismo-code-challenge/pkg/dtos"
)

type transaction struct{}

func (t *transaction) CreateTransaction(ctx context.Context, req *dtos.CreateTransaction) error {
	return nil
}
