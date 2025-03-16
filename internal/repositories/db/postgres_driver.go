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

func New(ctx context.Context) (*sql.DB, error) {
	conn, err := sql.Open("postgres", "")
	if err != nil {
		return nil, fmt.Errorf("db: %w", err)
	}

	return conn, nil
}
