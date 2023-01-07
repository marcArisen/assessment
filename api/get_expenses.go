package api

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/marcArisen/assessment/database"
)

func GetByIdExpenses(c echo.Context) error {

	exp, err := database.GetById(c.Param("id"))

	switch err {
	case sql.ErrNoRows:
		return c.JSON(http.StatusNotFound, "Expense doesn't exist")
	case nil:
		return c.JSON(http.StatusOK, exp)
	default:
		return c.JSON(http.StatusInternalServerError, err)
	}
}
