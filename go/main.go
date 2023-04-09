package main

import (
	"net/http"
	"strconv"

	"github.com/McCdama/govue/router"
	"github.com/labstack/echo"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type Pass struct {
	Orginal string `json:"orginal" xml:"orginal" form:"orginal" query:"orginal"`
	Hash    string `json:"hash" xml:"hash" form:"hash" query:"hash"`
}

type Bcrypt struct {
	Password string `json:"password" xml:"password" form:"password" query:"password"`
	Rounds   int    `json:"rounds" xml:"rounds" form:"rounds" query:"rounds"`
	Hash     []byte `json:"hash" xml:"hash" form:"hash" query:"hash"`
}

type IDs struct {
	Times int       `json:"times" xml:"times" form:"times" query:"times"`
	Uuid  uuid.UUID `json:"uuid" xml:"uuid" form:"uuid" query:"uuid"`
}

func main() {
	e := router.New()
	e.Logger.Fatal(e.Start(":8080"))
}

func gouuid(c echo.Context) error {

	numberRequested, _ := strconv.Atoi(c.FormValue("numberRequested"))

	myuuid := uuid.NewV4()

	u := &IDs{
		Times: numberRequested,
		Uuid:  myuuid,
	}

	return c.JSON(http.StatusOK, u)

}

func goBcrypt(c echo.Context) error {

	password := c.FormValue("password")
	roundsString := c.FormValue("rounds")
	rounds, _ := strconv.Atoi(roundsString)

	bcryptedPass, err := bcrypt.GenerateFromPassword([]byte(password), rounds)
	if err != nil {
		panic(err)
	}

	u := &Bcrypt{
		Password: password,
		Rounds:   rounds,
		Hash:     bcryptedPass,
	}

	return c.JSON(http.StatusOK, u)
}

func gohashPass(c echo.Context) error {
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
