package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"golang.org/x/crypto/bcrypt"
)

type Pass struct {
	Orginal string `json:"orginal" xml:"orginal"`
	Hash    string `json:"hash" xml:"hash"`
}

type Bcrypt struct {
	Password string `json:"password" xml:"password"`
	Rounds   int    `json:"rounds" xml:"rounds"`
	Hash     []byte `json:"hash" xml:"hash"`
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

	e.POST("/bcrypt", goBcrypt)

	//genBcrypt()

	//checkHash([]byte("$2a$10$PzOn.hp6aSI9s8jL7C66qOSQXOCKe8kXXsep6tXfhTCipTP9j2w0GHf."), []byte("mysecretpassword"))
	//checkHash([]byte("$2a$13$//NYiT3pjj6iAv7ZWPvzjO1pMh1ot/8RuK7GdYc81BKkt9R2Ca.Uy"), []byte("mysecretpassword"))

	e.Logger.Fatal(e.Start(":8080"))
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

// func checkHash(hashedPassword []byte, password []byte) {
// 	err := bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
// 	if err != nil {
// 		fmt.Println("Password does not match.")
// 	} else {
// 		fmt.Println("Password matches!")
// 	}
// }

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
