package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/marcArisen/assessment/database"
)

func GetAllExpenses(c echo.Context) error {

	exps, err := database.GetAllRecords()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, exps)
}
