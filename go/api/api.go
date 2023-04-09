package api

import (
	"github.com/McCdama/govue/handlers"
	"github.com/labstack/echo"
)

func UtilGroup(e *echo.Echo) {
	e.GET("/ping", handlers.Pong)
	//c.POST("/hash", gohashPass)
	//
	//c.POST("/bcrypt", goBcrypt)
	//
	//c.POST("/uuid", gouuid)
}
