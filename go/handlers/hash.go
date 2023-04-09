package handlers

import (
	"net/http"

	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

type Pass struct {
	Orginal string `json:"orginal" xml:"orginal" form:"orginal" query:"orginal"`
	Hash    string `json:"hash" xml:"hash" form:"hash" query:"hash"`
}

func Gohash(c echo.Context) error {
	p := new(Pass)
	// 👇 parses the body and "binds" it to the passed type
	if err := c.Bind(p); err != nil {
		return c.JSON(http.StatusBadRequest, "Please provid a payload..")
	}

	rawPassword := c.FormValue("rawPassword")

	if rawPassword == "" || len(rawPassword) < 8 {
		return c.JSON(http.StatusBadRequest, "Please provide a password that is at least 8 characters long")
	}

	hash, err := hashPassword(rawPassword)
	if err != nil {
		return c.JSON(http.StatusExpectationFailed, "Something went wrong..")
	}

	p = &Pass{
		Orginal: rawPassword,
		Hash:    hash,
	}

	return c.JSON(http.StatusOK, p)
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
