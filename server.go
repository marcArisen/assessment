package main

import (
	"fmt"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	fmt.Println("Please use server.go for main file")
	fmt.Println("start at port:", port)
	database.init()

	// db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	// if err != nil {
	// 	log.Fatal("Connect to database error", err)
	// }
	// createTb := `
	// CREATE TABLE IF NOT EXISTS expenses (
	// 	id SERIAL PRIMARY KEY,
	// 	title TEXT,
	// 	amount FLOAT,
	// 	note TEXT,
	// 	tags TEXT[]
	// );
	// `
	// _, err := db.Exec(createTb)

}
