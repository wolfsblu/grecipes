package db

import (
	"database/sql"
	"log"
	_ "modernc.org/sqlite"
)

var Query *Queries

func init() {
	con, err := sql.Open("sqlite", "tmp/db.sqlite")
	if err != nil {
		log.Fatalln("failed to open db:", err)
	}
	Query = New(con)
}
