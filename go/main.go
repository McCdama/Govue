package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"golang.org/x/crypto/bcrypt"
)

type Pass struct {
	Orginal string `json:"orginal" xml:"orginal"`
	Hash    string `json:"hash" xml:"hash"`
}

func main() {

	e := echo.New()
	// e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
	// 	Skipper: func(c echo.Context) bool {
	// 		if strings.HasPrefix(c.Request().Host, "localhost") {
	// 			return true
	// 		}
	// 		return false
	// 	},
	// }))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong!")
	})
	e.POST("/hash", generateHashedPass)
	e.Logger.Fatal(e.Start(":8080"))
}

func generateHashedPass(c echo.Context) error {
	rawPassword := c.FormValue("rawPassword")
	hash, _ := hashPassword(rawPassword)

	u := &Pass{
		Orginal: rawPassword,
		Hash:    hash,
	}

	return c.JSON(http.StatusOK, u)

	//return c.HTML(http.StatusOK, "<b>Thank you! "+hash+"</b>")
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
