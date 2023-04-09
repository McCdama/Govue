package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

func Pong(c echo.Context) error {

	return c.JSON(http.StatusOK, "pong!")
}
