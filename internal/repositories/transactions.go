package repositories

import (
	"context"

	"github.com/ppcamp/go-pismo-code-challenge/internal/models"
	"github.com/ppcamp/go-pismo-code-challenge/internal/repositories/db"
	errutils "github.com/ppcamp/go-pismo-code-challenge/pkg/utils/errors"
)

//go:generate -command mockgen go tool mockgen

//go:generate mockgen -destination=mock/transactions.go . Transactions
type Transactions interface {
	Create(ctx context.Context, conn db.Driver, data *models.Transaction) error
}

type implTransactions struct{}

func NewTransactions() Transactions { return &implTransactions{} }

func (t *implTransactions) Create(ctx context.Context, conn db.Driver, data *models.Transaction) error {
	const query = `
		INSERT INTO pismo.transactions (account_id, operation_type_id, amount) 
		VALUES ($1, $2, $3)`

	err := conn.Exec(ctx, query, data.AccountId, data.OperationId, data.Amount)
	if err != nil {
		return errutils.Error{Base: db.ErrDriverError, Wrapped: err}
	}

	return nil
}
