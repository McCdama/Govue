package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

type Bcrypt struct {
	Password string `json:"password" xml:"password" form:"password" query:"password"`
	Rounds   int    `json:"rounds" xml:"rounds" form:"rounds" query:"rounds"`
	Hash     []byte `json:"hash" xml:"hash" form:"hash" query:"hash"`
}

func GoBcrypt(c echo.Context) error {

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
