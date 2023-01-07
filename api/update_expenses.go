package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/marcArisen/assessment/database"
)

func UpdateExpenses(c echo.Context) error {

	exp, err := database.GetById(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	err = c.Bind(&exp)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	exp, err = database.Update(exp)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, exp)
}
