package repositories

import (
	"context"
	"database/sql"
	"errors"

	"github.com/ppcamp/go-pismo-code-challenge/internal/models"
	"github.com/ppcamp/go-pismo-code-challenge/internal/repositories/db"
	errutils "github.com/ppcamp/go-pismo-code-challenge/pkg/utils/errors"
)

type Account interface {
	Create(ctx context.Context, account *models.Account) error
	Get(ctx context.Context, id int64) (*models.Account, error)
}

type implAccount struct{ conn db.Driver }

func NewAccount(conn db.Driver) Account { return &implAccount{conn} }

func (t *implAccount) Create(ctx context.Context, data *models.Account) error {
	const query = `INSERT INTO accounts(document_number) VALUES ($1)`

	err := t.conn.Exec(ctx, query, data.DocumentNumber)
	if err != nil {
		return errutils.Error{Base: db.ErrDriverError, Wrapped: err}
	}

	return nil
}

func (t *implAccount) Get(ctx context.Context, id int64) (*models.Account, error) {
	const query = `SELECT id, document_number FROM accounts WHERE id = $1`

	row := t.conn.QueryRow(ctx, query, id)

	var acct models.Account
	err := row.Scan(&acct.Id, &acct.DocumentNumber)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errutils.Error{Base: db.ErrNotFound, Wrapped: err}
		}

		return nil, errutils.Error{Base: db.ErrDriverError, Wrapped: err}
	}

	return &acct, nil
}
