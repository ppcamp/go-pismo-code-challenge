package repositories

import (
	"context"

	"github.com/ppcamp/go-pismo-code-challenge/internal/models"
	"github.com/ppcamp/go-pismo-code-challenge/internal/repositories/db"
	errutils "github.com/ppcamp/go-pismo-code-challenge/pkg/utils/errors"
)

type Transactions interface {
	Create(ctx context.Context, data *models.Transaction) error
}

type implTransactions struct{ conn db.Driver }

func NewTransactions(conn db.Driver) Transactions { return &implTransactions{conn} }

func (t *implTransactions) Create(ctx context.Context, data *models.Transaction) error {
	const query = `
		INSERT INTO transactions (account_id, operation_type_id, amount) 
		VALUES ($1, $2, $3)`

	_, err := t.conn.ExecContext(ctx, query, data.AccountId, data.OperationId, data.Amount)
	if err != nil {
		return errutils.Error{Base: db.ErrDriverError, Wrapped: err}
	}
	return nil
}
