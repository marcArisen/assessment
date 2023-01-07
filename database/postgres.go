package database

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
	"github.com/marcArisen/assessment/model"
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

func CreateExpenses(c echo.Context) error {
	var exp model.Expenses
	var err error
	err = c.Bind(&exp)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	exp, err = insert(exp)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusCreated, exp)
}

func insert(exp model.Expenses) (model.Expenses, error) {

	insertStatement := `
		INSERT INTO expenses(title,amount,note,tags) values($1,$2,$3,$4) RETURNING *
	`

	row := db.QueryRow(insertStatement, exp.Title, exp.Amount, exp.Note, pq.Array(exp.Tags))
	err := row.Scan(&exp.Id, &exp.Title, &exp.Amount, &exp.Note, &exp.Tags)

	if err != nil {
		return model.Expenses{}, err
	}

	return exp, nil
}

func init() {

	var err error

	db, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))

	if err != nil {
		log.Fatal("Cannot Connect to The Database", err)
	}

	createTable(db)

}
