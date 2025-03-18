package db

import (
	"context"
	"fmt"
)

// Transaction is a decorator that creates a transaction and, accordding to the
// fn return, commit or rollback the db change.
func Transaction(ctx context.Context, conn DB, cb func(ctx context.Context, db DriverTransaction) error) (err error) {
	tx, err := conn.BeginTx(ctx)
	if err != nil {
		return fmt.Errorf("%w: fail to create transaction: %w", ErrDriverError, err)
	}

	err = cb(ctx, tx)
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panick(ed) in the Transaction Decorator: %v", r)
		}

		if err != nil {
			errTx := tx.Rollback(ctx)
			if errTx != nil {
				err = fmt.Errorf("fail to rollback: %w: %w", errTx, err)
			}
			return
		}

		err = tx.Commit(ctx)
	}()

	return
}
