package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

type Params struct {
	Driver         string
	Host           string
	Port           int
	User, Password string
}

func (p Params) ConnStr() string {
	return fmt.Sprintf("%s://%s:%s@%s:%d", p.Driver, p.User, p.Password, p.Host, p.Port)
}

// New implements a postgres driver with connection pooling.
//
// This is concurrent safe pooling
func New(ctx context.Context, params Params) (DB, error) {
	logrus.WithField("connection_string", params.ConnStr()).
		Info("Creating connection")
	// The main reason to use pgx insteadk of the default sql driver is due to
	// connection pooling, which will guarantee that we don't open too many
	// connections, neither keep them alive when no service is using.
	conn, err := pgxpool.New(ctx, params.ConnStr())
	if err != nil {
		return nil, fmt.Errorf("db: %w", err)
	}

	return &implDriver{conn, &implBasicDriver{conn}}, nil
}

type implDriver struct {
	conn *pgxpool.Pool
	*implBasicDriver
}

func (i *implDriver) Ping(ctx context.Context) error { return i.conn.Ping(ctx) }

func (i *implDriver) Close(ctx context.Context) error {
	i.conn.Close()
	return nil
}

// BeginTx initializes a new sql transaction.
// This, however, should not be handled manually. Check for
// [github.com/ppcamp/go-pismo-code-challenge/internal/repositories/db.]
func (i *implDriver) BeginTx(ctx context.Context) (DriverTransaction, error) {
	tx, err := i.conn.Begin(ctx)
	if err != nil {
		return nil, err
	}

	return &implTransactionDriver{tx, &implBasicDriver{i.conn}}, nil
}

type implTransactionDriver struct {
	Tx pgx.Tx
	*implBasicDriver
}

func (i *implTransactionDriver) Commit(ctx context.Context) error   { return i.Tx.Commit(ctx) }
func (i *implTransactionDriver) Rollback(ctx context.Context) error { return i.Tx.Rollback(ctx) }

type implBasicDriver struct{ conn *pgxpool.Pool }

func (i *implBasicDriver) Exec(ctx context.Context, query string, args ...any) error {
	_, err := i.conn.Exec(ctx, query, args...)
	return err
}

func (i *implBasicDriver) Query(ctx context.Context, query string, args ...any) (Rows, error) {
	return i.conn.Query(ctx, query, args...)
}

func (i *implBasicDriver) QueryRow(ctx context.Context, query string, args ...any) Row {
	return i.conn.QueryRow(ctx, query, args...)
}
