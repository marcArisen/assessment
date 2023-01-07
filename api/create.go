package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/marcArisen/assessment/database"
	"github.com/marcArisen/assessment/model"
)

func CreateExpenses(c echo.Context) error {
	var exp model.Expenses
	var err error
	err = c.Bind(&exp)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	exp, err = database.Insert(exp)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusCreated, exp)
}
