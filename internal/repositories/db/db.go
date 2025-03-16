package db

type Driver interface {
	StartTransaction()

	Ping() error
	Close() error
}
