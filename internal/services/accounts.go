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
