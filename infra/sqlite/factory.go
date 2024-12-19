package sqlite

import (
	"database/sql"
	"fmt"
	"github.com/wolfsblu/go-chef/infra/env"
)

type Store struct {
	db   *sql.DB
	path string
	q    *Queries
	qtx  *Queries
	tx   *sql.Tx
}

func NewSqliteStore() (*Store, error) {
	dbPath := env.MustGet("DB_PATH")
	con, err := connect(dbPath)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %w", err)
	}
	return &Store{db: con, q: New(con), path: dbPath}, nil
}

func connect(path string) (*sql.DB, error) {
	con, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, err
	}
	return con, nil
}
