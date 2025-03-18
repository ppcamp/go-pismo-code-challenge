package services

import (
	"context"
	"fmt"

	"github.com/ppcamp/go-pismo-code-challenge/internal/models"
	"github.com/ppcamp/go-pismo-code-challenge/internal/repositories"
	"github.com/ppcamp/go-pismo-code-challenge/internal/repositories/db"
	"github.com/ppcamp/go-pismo-code-challenge/pkg/dtos"
)

type Account interface {
	Create(ctx context.Context, req *dtos.CreateAccount) error
	Get(ctx context.Context, id int64) (*dtos.Account, error)
}

type implAccount struct{ conn db.DB }

func NewAccountService(dbconn db.DB) Account { return &implAccount{dbconn} }

func (t *implAccount) Create(ctx context.Context, req *dtos.CreateAccount) error {
	pl := &models.Account{DocumentNumber: req.DocumentNumber}

	return db.Transaction(ctx, t.conn, func(ctx context.Context, db db.DriverTransaction) error {
		repo := repositories.NewAccount(db)

		err := repo.Create(ctx, pl)
		if err != nil {
			return fmt.Errorf("fail to create account: %w", err)
		}
		return nil
	})
}

func (t *implAccount) Get(ctx context.Context, id int64) (*dtos.Account, error) {
	repo := repositories.NewAccount(t.conn)

	acct, err := repo.Get(ctx, id)
	if err == nil {
		return &dtos.Account{Id: acct.Id, DocumentNumber: acct.DocumentNumber}, nil
	}
	return nil, fmt.Errorf("fail to get account: %w", err)

}
