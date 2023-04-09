package api

import (
	"github.com/McCdama/govue/handlers"
	"github.com/labstack/echo"
)

func UtilGroup(e *echo.Echo) {
	e.GET("/ping", handlers.Pong)
	e.POST("/hash", handlers.Gohash)
	e.POST("/bcrypt", handlers.GoBcrypt)
	e.POST("/uuid", handlers.Gouuid)
}
