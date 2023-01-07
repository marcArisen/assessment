package database

import (
	"database/sql"
	"log"
	"os"

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

func Insert(exp model.Expenses) (model.Expenses, error) {

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

func GetById(id string) (model.Expenses, error) {

	queryStatement := `
	SELECT * FROM expenses WHERE id = $1
	`

	st, err := db.Prepare(queryStatement)

	if err != nil {
		return model.Expenses{}, err
	}

	exp := model.Expenses{}
	err = st.QueryRow(id).Scan(&exp.Id, &exp.Title, &exp.Amount, &exp.Note, &exp.Tags)

	if err != nil {
		return model.Expenses{}, err
	}

	return exp, nil

}

func GetAllRecords() ([]model.Expenses, error) {

	queryStatement := `
	SELECT * FROM expenses
	`

	st, err := db.Prepare(queryStatement)

	if err != nil {
		return nil, err
	}

	rows, err2 := st.Query()
	if err2 != nil {
		return nil, err
	}

	exps := []model.Expenses{}

	for rows.Next() {

		exp := model.Expenses{}

		err = rows.Scan(&exp.Id, &exp.Title, &exp.Amount, &exp.Note, &exp.Tags)

		if err != nil {
			return exps, err
		}

		exps = append(exps, exp)
	}

	return exps, nil

}

func Update(exp model.Expenses) (model.Expenses, error) {

	updateStatement := `
	UPDATE expenses SET title=$2 ,amount=$3 ,note=$4 ,tags=$5 WHERE id=$1 RETURNING *
	`

	err := db.QueryRow(updateStatement, exp.Id, exp.Title, exp.Amount, exp.Note, exp.Tags).Scan(&exp.Id, &exp.Title, &exp.Amount, &exp.Note, &exp.Tags)

	if err != nil {
		return exp, err
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
