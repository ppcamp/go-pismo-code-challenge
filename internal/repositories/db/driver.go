package db

import (
	"context"
	"database/sql"
)

type Driver interface {
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
}

type DriverTransaction interface {
	Commit() error
	Rollback() error

	Driver
}

type DB interface {
	// BeginTx initializes a new sql transaction.
	//
	// This, however, should not be handled manually. Check for Transaction
	// decorator.
	BeginTx(ctx context.Context) (DriverTransaction, error)

	Close() error
	Ping() error

	Driver
}
