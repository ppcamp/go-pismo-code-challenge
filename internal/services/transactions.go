package services

import (
	"context"
	"fmt"

	"github.com/ppcamp/go-pismo-code-challenge/internal/models"
	"github.com/ppcamp/go-pismo-code-challenge/internal/repositories"
	"github.com/ppcamp/go-pismo-code-challenge/internal/repositories/db"
	"github.com/ppcamp/go-pismo-code-challenge/pkg/dtos"
	"github.com/sirupsen/logrus"
)

type Transaction interface {
	Create(ctx context.Context, req *dtos.CreateTransaction) error
}

type implTransaction struct{ conn db.DB }

func NewTransactionService(dbconn db.DB) Transaction { return &implTransaction{dbconn} }

func (t *implTransaction) Create(ctx context.Context, req *dtos.CreateTransaction) error {
	log := logrus.WithContext(ctx).WithField("payload", req)

	pl := &models.Transaction{
		AccountId:   req.AccountId,
		OperationId: req.OperationTypeId,
		Amount:      req.Amount,
	}

	log.Info("trying to create transactions")
	return db.Transaction(ctx, t.conn, func(ctx context.Context, db db.DriverTransaction) error {
		repo := repositories.NewTransactions(db)

		log.Info("inserting transaction into database")
		err := repo.Create(ctx, pl)
		if err != nil {
			log.WithError(err).Error("fail to create transaction")
			return fmt.Errorf("fail to create account: %w", err)
		}
		log.Info("transaction created successfuly")

		return nil
	})
}
