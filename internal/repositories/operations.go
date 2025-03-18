package repositories

import (
	"database/sql"

	"github.com/ppcamp/go-pismo-code-challenge/internal/repositories/db"
)

type Operations interface{}

type implOperations struct {
	conn *sql.DB
}

func NewOperations(conn db.Driver) Operations { return &implOperations{} }
