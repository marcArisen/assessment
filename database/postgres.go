package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

var db *sql.DB

func createTable(db *sql.DB) {

	initTable := `
	CREATE TABLE IF NOT EXISTS expenses (
		id SERIAL PRIMARY KEY,
		title TEXT,
		amount FLOAT,
		note TEXT,
		tags TEXT[]
	);
	`
	_, err := db.Exec(initTable)

	if err != nil {
		log.Fatal("Having a Problem Creating new Table", err)
	}

}

func init() {

	fmt.Printf("GOin")

	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))

	if err != nil {
		log.Fatal("Cannot Connect to The Database", err)
	}

	createTable(db)

}
