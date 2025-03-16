package repositories

import (
	"database/sql"
)

type Operations interface{}

type implOperations struct {
	conn *sql.DB
}

func NewOperations(conn *sql.DB) Operations { return &implOperations{} }
