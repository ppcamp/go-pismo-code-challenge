package services

import (
	"context"
	"fmt"

	"github.com/ppcamp/go-pismo-code-challenge/internal/models"
	"github.com/ppcamp/go-pismo-code-challenge/internal/repositories"
	"github.com/ppcamp/go-pismo-code-challenge/internal/repositories/db"
	"github.com/ppcamp/go-pismo-code-challenge/pkg/dtos"
)

type Transaction interface {
	Create(ctx context.Context, req *dtos.CreateTransaction) error
}

type implTransaction struct{ conn db.DB }

func (t *implTransaction) Create(ctx context.Context, req *dtos.CreateTransaction) error {
	pl := &models.Transaction{
		AccountId:   req.AccountId,
		OperationId: req.OperationTypeId,
		Amount:      req.Amount,
	}

	return db.Transaction(ctx, t.conn, func(ctx context.Context, db db.DriverTransaction) error {
		repo := repositories.NewTransactions(db)

		err := repo.Create(ctx, pl)
		if err != nil {
			return fmt.Errorf("fail to create account: %w", err)
		}
		return nil
	})
}
