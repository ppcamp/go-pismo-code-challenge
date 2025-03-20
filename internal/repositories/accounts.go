package repositories

import (
	"context"
	"database/sql"
	"errors"

	"github.com/ppcamp/go-pismo-code-challenge/internal/models"
	"github.com/ppcamp/go-pismo-code-challenge/internal/repositories/db"
	errutils "github.com/ppcamp/go-pismo-code-challenge/pkg/utils/errors"
)

//go:generate -command mockgen go tool mockgen

//go:generate mockgen -destination=mock/accounts.go . Account
type Account interface {
	// Create an Account
	//
	// The connection is passed, so we can handle transactions outside this
	// module, and therefore, join several queries at a single transacion,
	// rollback if necessary.
	Create(ctx context.Context, conn db.Driver, account *models.Account) error
	Get(ctx context.Context, conn db.Driver, id int64) (*models.Account, error)
}

type implAccount struct{}

func NewAccount() Account { return &implAccount{} }

func (t *implAccount) Create(ctx context.Context, conn db.Driver, data *models.Account) error {
	const query = `INSERT INTO pismo.accounts(document_number) VALUES ($1)`

	err := conn.Exec(ctx, query, data.DocumentNumber)
	if err != nil {
		return errutils.Error{Base: db.ErrDriverError, Wrapped: err}
	}

	return nil
}

func (t *implAccount) Get(ctx context.Context, conn db.Driver, id int64) (*models.Account, error) {
	const query = `SELECT id, document_number FROM pismo.accounts WHERE id = $1`

	row := conn.QueryRow(ctx, query, id)

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
