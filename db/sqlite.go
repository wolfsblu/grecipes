package db

import (
	"database/sql"
	_ "modernc.org/sqlite"
)

func Connect(path string) (*Queries, error) {
	con, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, err
	}
	return New(con), nil
}
