package sqlite

import (
	"database/sql"
	"fmt"
)

func NewSqliteStore(dbPath string) (*Store, error) {
	query, err := connect(dbPath)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %w", err)
	}
	return &Store{db: query, path: dbPath}, nil
}

func connect(path string) (*Queries, error) {
	con, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, err
	}
	return New(con), nil
}
