package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Params struct {
	Driver         string
	Host           string
	Port           int
	User, Password string
}

func (p *Params) ConnStr() string {
	return fmt.Sprintf("%s:%d@%s:%s", p.Host, p.Port, p.User, p.Password)
}

func New(params Params) (DB, error) {
	conn, err := sql.Open(params.Driver, params.ConnStr())
	if err != nil {
		return nil, fmt.Errorf("db: %w", err)
	}

	return &implDriver{conn, &implBasicDriver{conn}}, nil
}

type implDriver struct {
	conn *sql.DB
	*implBasicDriver
}

func (i *implDriver) Ping() error  { return i.conn.Ping() }
func (i *implDriver) Close() error { return i.conn.Close() }

// BeginTx initializes a new sql transaction.
// This, however, should not be handled manually. Check for
// [github.com/ppcamp/go-pismo-code-challenge/internal/repositories/db.]
func (i *implDriver) BeginTx(ctx context.Context) (DriverTransaction, error) {
	tx, err := i.conn.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	return &implTransactionDriver{tx}, nil
}

type implTransactionDriver struct{ *sql.Tx }

func (i *implTransactionDriver) Commit() error   { return i.Tx.Commit() }
func (i *implTransactionDriver) Rollback() error { return i.Tx.Rollback() }

type implBasicDriver struct{ conn *sql.DB }

func (i *implBasicDriver) ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error) {
	return i.conn.ExecContext(ctx, query, args...)
}

func (i *implBasicDriver) QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error) {
	return i.conn.QueryContext(ctx, query, args...)
}

func (i *implBasicDriver) QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row {
	return i.conn.QueryRowContext(ctx, query, args...)
}
