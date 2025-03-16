package repositories

import (
	"context"
	"database/sql"

	"github.com/ppcamp/go-pismo-code-challenge/internal/models"
)

type Transactions interface {
	Create(ctx context.Context, data models.Transaction) error
}

type implTransactions struct {
	conn *sql.DB
}

func NewTransactions() Transactions { return &implTransactions{} }

func (t *implTransactions) Create(ctx context.Context, data models.Transaction) error {
	const query = `
		INSERT INTO transactions (account_id, operation_type_id, amount) 
		VALUES ($1, $2, $3)`

	_, err := t.conn.ExecContext(ctx, query, data.AccountId, data.OperationId, data.Amount)
	if err != nil {
		return err
	}

	return nil
}
