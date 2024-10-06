package db

import (
	"database/sql"
	"log"
	"os"
	_ "modernc.org/sqlite"
)

var Query *Queries

func init() {
	dbPath, ok := os.LookupEnv("DB_PATH")
	if !ok {
		log.Fatalln("variable DB_PATH needs to be set to a valid sqlite database path")
	}
	con, err := sql.Open("sqlite", dbPath)
	if err != nil {
		log.Fatalln("failed to open db:", err)
	}
	Query = New(con)
}
