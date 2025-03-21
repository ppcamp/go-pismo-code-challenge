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

type Account interface {
	Create(ctx context.Context, req *dtos.CreateAccount) error
	Get(ctx context.Context, id int64) (*dtos.Account, error)

	SetLimit(ctx context.Context, id int64, newLimit float64) error
	GetAccountLimits(ctx context.Context, id int64) (*dtos.AccountLimits, error)

	RemoveLimit(ctx context.Context, id int64, toRemove float64) error
	// AddLimit Usually for payments
	AddLimit(ctx context.Context, id int64, toAdd float64) error
}

type implAccount struct {
	conn db.DB
	repo repositories.Account
}

func NewAccountService(dbconn db.DB, repo repositories.Account) Account {
	return &implAccount{dbconn, repo}
}

func (t *implAccount) Create(ctx context.Context, req *dtos.CreateAccount) error {
	log := logrus.WithContext(ctx).WithField("payload", req)

	pl := &models.Account{DocumentNumber: req.DocumentNumber}

	log.Info("creating account")
	return db.Transaction(ctx, t.conn, func(ctx context.Context, db db.DriverTransaction) error {
		log.Info("inserting account into db")

		err := t.repo.Create(ctx, db, pl)
		if err != nil {
			log.WithError(err).Error("fail to create account")
			return fmt.Errorf("fail to create account: %w", err)
		}

		log.Info("account created successfuly")
		return nil
	})
}

func (t *implAccount) Get(ctx context.Context, id int64) (*dtos.Account, error) {
	log := logrus.WithContext(ctx).WithField("account_id", id)

	log.Info("fetching account")
	acct, err := t.repo.Get(ctx, t.conn, id)
	if err == nil {
		log.Info("account fetched successfuly")
		return &dtos.Account{Id: acct.Id, DocumentNumber: acct.DocumentNumber}, nil
	}

	log.WithError(err).Error("fail to get account")
	return nil, fmt.Errorf("fail to get account: %w", err)

}

func (t *implAccount) GetAccountLimits(ctx context.Context, id int64) (*dtos.AccountLimits, error) {
	log := logrus.WithContext(ctx).WithField("account_id", id)

	log.Info("fetching account limits")
	acct, err := t.repo.GetLimits(ctx, t.conn, id)
	if err == nil {
		log.Info("account fetched successfuly")
		return &dtos.AccountLimits{AvailableLimit: acct.AvailableLimit, CurrentLimit: acct.CurrentLimit}, nil
	}

	log.WithError(err).Error("fail to get limits for account")
	return nil, fmt.Errorf("fail to get account: %w", err)

}

func (t *implAccount) RemoveLimit(ctx context.Context, id int64, toRemove float64) error {
	log := logrus.WithContext(ctx).WithField("account_id", id)

	log.Info("fetching account limits")
	acct, err := t.repo.GetLimits(ctx, t.conn, id)
	if err != nil {
		return fmt.Errorf("fail to get account limits: %w", err)
	}

	nl := acct.CurrentLimit - toRemove
	if nl < 0 {
		return fmt.Errorf("cannot change the limit due to max acct limit")
	}
	log.WithError(err).Error("fail to get limits for account")

	return db.Transaction(ctx, t.conn, func(ctx context.Context, db db.DriverTransaction) error {
		return t.repo.UpdateLimits(ctx, db, id, &models.AccountLimit{
			AvailableLimit: acct.AvailableLimit,
			CurrentLimit:   nl,
		})
	})
}

func (t *implAccount) AddLimit(ctx context.Context, id int64, toAdd float64) error {
	log := logrus.WithContext(ctx).WithField("account_id", id)

	log.Info("fetching account limits")
	acct, err := t.repo.GetLimits(ctx, t.conn, id)
	if err != nil {
		return fmt.Errorf("fail to get account limits: %w", err)
	}

	nl := acct.CurrentLimit + toAdd

	return db.Transaction(ctx, t.conn, func(ctx context.Context, db db.DriverTransaction) error {
		return t.repo.UpdateLimits(ctx, db, id, &models.AccountLimit{
			AvailableLimit: acct.AvailableLimit,
			CurrentLimit:   nl,
		})
	})
}

func (t *implAccount) SetLimit(ctx context.Context, id int64, newLimit float64) error {
	log := logrus.WithContext(ctx).WithField("account_id", id)

	log.Info("fetching account limits")
	acct, err := t.repo.GetLimits(ctx, t.conn, id)
	if err != nil {
		return fmt.Errorf("fail to get account limits: %w", err)
	}

	return db.Transaction(ctx, t.conn, func(ctx context.Context, db db.DriverTransaction) error {
		return t.repo.UpdateLimits(ctx, db, id, &models.AccountLimit{
			AvailableLimit: newLimit,
			CurrentLimit:   acct.CurrentLimit,
		})
	})
}
