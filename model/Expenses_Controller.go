package model

// import (
// 	"database/sql"
// 	"net/http"

// 	"github.com/labstack/echo/v4"
// )

// func CreateExpenses(db *sql.DB, c echo.Context) error {
// 	var exp Expenses
// 	var err error
// 	err = c.Bind(&exp)

// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, err.Error())
// 	}

// 	exp, err = database.Insert(exp)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, err)
// 	}
// 	return c.JSON(http.StatusCreated, exp)
// }
