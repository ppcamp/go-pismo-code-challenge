package services

import (
	"context"
	"fmt"

	"github.com/ppcamp/go-pismo-code-challenge/internal/models"
	"github.com/ppcamp/go-pismo-code-challenge/internal/repositories"
	"github.com/ppcamp/go-pismo-code-challenge/internal/repositories/db"
	"github.com/ppcamp/go-pismo-code-challenge/pkg/dtos"
	"github.com/ppcamp/go-pismo-code-challenge/pkg/enums"
	"github.com/sirupsen/logrus"
)

type Transaction interface {
	Create(ctx context.Context, req *dtos.CreateTransaction) error
}

type implTransaction struct {
	conn db.DB
	repo repositories.Transactions
	acct Account
}

func NewTransactionService(dbconn db.DB, repo repositories.Transactions, acct Account) Transaction {
	return &implTransaction{dbconn, repo, acct}
}

func (t *implTransaction) Create(ctx context.Context, req *dtos.CreateTransaction) error {
	log := logrus.WithContext(ctx).WithField("payload", req)

	pl := &models.Transaction{
		AccountId:   req.AccountId,
		OperationId: req.OperationTypeId,
		Amount:      req.Amount,
	}

	limits, err := t.acct.GetAccountLimits(ctx, req.AccountId)
	if err != nil {
		return fmt.Errorf("fail to get: %w", err)
	}

	return db.Transaction(ctx, t.conn, func(ctx context.Context, db db.DriverTransaction) error {
		var err error

		switch req.OperationTypeId {
		case enums.OpCreditVoucher, enums.OpWithdrawl:
			err = t.doPayment(ctx, pl, log)
		default:
			err = t.doPurchase(ctx, limits, pl, log)
		}

		if err != nil {
			return fmt.Errorf("some error occurred in account service: %w", err)
		}

		log.Info("inserting transaction into database")
		err = t.repo.Create(ctx, db, pl)
		if err != nil {
			log.WithError(err).Error("fail to create transaction")
			return fmt.Errorf("fail to create account: %w", err)
		}
		log.Info("transaction created successfuly")

		return nil

	})
}

func (t *implTransaction) doPayment(ctx context.Context, pl *models.Transaction, log *logrus.Entry) error {
	err := t.acct.AddLimit(ctx, pl.AccountId, pl.Amount)
	if err != nil {
		log.WithError(err).Error("fail to add limit to account")
		return fmt.Errorf("cannot reduce account limit: %w", err)
	}
	return nil
}

func (t *implTransaction) doPurchase(ctx context.Context, limits *dtos.AccountLimits, pl *models.Transaction, log *logrus.Entry) error {
	if pl.Amount > limits.CurrentLimit {
		log.WithField("available", limits.CurrentLimit).Warn("fail to perform transaction to insuficient limit")
		return fmt.Errorf("transaction failed due to insuficient limit")
	}

	err := t.acct.RemoveLimit(ctx, pl.AccountId, pl.Amount)
	if err != nil {
		return fmt.Errorf("cannot reduce account limit: %w", err)
	}
	return nil
}
