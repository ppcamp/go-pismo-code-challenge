package repositories

import (
	"context"
	"database/sql"

	"github.com/ppcamp/go-pismo-code-challenge/internal/models"
)

type Account interface {
	Create(context.Context, models.Account) error
	Get(context.Context, int64) (*models.Account, error)
}

type implAccount struct {
	conn *sql.DB
}

func NewAccount(conn *sql.DB) Account { return &implAccount{} }

func (t *implAccount) Create(ctx context.Context, data models.Account) error {
	const query = `INSERT INTO accounts(document_number) VALUES ($1)`

	_, err := t.conn.ExecContext(ctx, query, data.DocumentNumber)
	if err != nil {
		return err
	}

	return nil
}

func (t *implAccount) Get(ctx context.Context, id int64) (*models.Account, error) {
	const query = `SELECT id, document_number FROM accounts WHERE id = $1`

	row := t.conn.QueryRowContext(ctx, query, id)

	var acct models.Account
	err := row.Scan(&acct.Id, &acct.DocumentNumber)

	return &acct, err
}
