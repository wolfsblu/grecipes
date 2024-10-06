package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var Query *Queries

func init() {
	con, err := sql.Open("sqlite3", "tmp/db.sqlite")
	if err != nil {
		log.Fatalln("failed to open db:", err)
	}
	Query = New(con)
}
