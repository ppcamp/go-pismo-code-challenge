package handlers

import (
	"context"
	"fmt"

	"github.com/ppcamp/go-pismo-code-challenge/internal/repositories"
	"github.com/ppcamp/go-pismo-code-challenge/pkg/dtos"
)

type Account interface {
	CreateAccount(ctx context.Context, req *dtos.CreateAccount) error
	GetAccount(ctx context.Context, id int64) (*dtos.Account, error)
}

type account struct {
	Repo repositories.Account
}

func NewAccountService() Account { return &account{} }

func (t *account) CreateAccount(ctx context.Context, req *dtos.CreateAccount) error {
	err := t.Repo.Create()
	if err != nil {
		return fmt.Errorf("fail to create account: %w", err)
	}

	return nil
}

func (t *account) GetAccount(ctx context.Context, id int64) (*dtos.Account, error) {
	return nil, nil
}
