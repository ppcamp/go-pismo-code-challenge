package repositories

import (
	"github.com/ppcamp/go-pismo-code-challenge/internal/repositories/db"
)

type Operations interface{}

type implOperations struct{ conn db.DB }

func NewOperations(conn db.Driver) Operations { return &implOperations{} }
