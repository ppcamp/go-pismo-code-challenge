package handlers

import (
	"context"

	"github.com/ppcamp/go-pismo-code-challenge/pkg/dtos"
)

type account struct{}

func (t *account) CreateAccount(ctx context.Context, req *dtos.CreateAccount) error {
	return nil
}

func (t *account) GetAccount(ctx context.Context, id int64) (*dtos.Account, error) {
	return nil, nil
}
