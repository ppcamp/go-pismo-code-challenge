package db

import (
	"context"
)

//go:generate -command mockgen go tool mockgen

//go:generate mockgen -destination=mock/row.go . Row,Rows
//go:generate mockgen -destination=mock/driver.go . Driver,DriverTransaction,DB

type Row interface {
	// Scan works the same as Rows. with the following exceptions. If no
	// rows were found it returns ErrNoRows. If multiple rows are returned it
	// ignores all but the first.
	Scan(dest ...any) error
}

type Rows interface {
	// Close closes the rows, making the connection ready for use again. It is safe
	// to call Close after rows is already closed.
	Close()

	// Err returns any error that occurred while reading. Err must only be called after the Rows is closed (either by
	// calling Close or by Next returning false). If it is called early it may return nil even if there was an error
	// executing the query.
	Err() error

	// Next prepares the next row for reading. It returns true if there is another
	// row and false if no more rows are available or a fatal error has occurred.
	// It automatically closes rows when all rows are read.
	//
	// Callers should check rows.Err() after rows.Next() returns false to detect
	// whether result-set reading ended prematurely due to an error. See
	// Conn.Query for details.
	//
	// For simpler error handling, consider using the higher-level pgx v5
	// CollectRows() and ForEachRow() helpers instead.
	Next() bool

	Row
}

// Drivers interfaces are just wrappers, since the default database/sql is a
// type, and I plan to keep it detached as possible to different drivers. And
// by doing so, if I choose to change the database, I just need to reimplement
// the required driver wrapper to match the defined interfaces here.
type Driver interface {
	Exec(ctx context.Context, query string, args ...any) error
	Query(ctx context.Context, query string, args ...any) (Rows, error)
	QueryRow(ctx context.Context, query string, args ...any) Row
}

type DriverTransaction interface {
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error

	Driver
}

type DB interface {
	// BeginTx initializes a new sql transaction.
	//
	// This, however, should not be handled manually. Check for Transaction
	// decorator.
	BeginTx(ctx context.Context) (DriverTransaction, error)

	Close(ctx context.Context) error
	Ping(ctx context.Context) error

	Driver
}
