package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
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
	e.POST("/hash", gohashPass)

	e.POST("/bcrypt", goBcrypt)

	e.POST("/uuid", gouuid)

	//genBcrypt()

	//checkHash([]byte("$2a$10$PzOn.hp6aSI9s8jL7C66qOSQXOCKe8kXXsep6tXfhTCipTP9j2w0GHf."), []byte("mysecretpassword"))
	//checkHash([]byte("$2a$13$//NYiT3pjj6iAv7ZWPvzjO1pMh1ot/8RuK7GdYc81BKkt9R2Ca.Uy"), []byte("mysecretpassword"))

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

// func checkHash(hashedPassword []byte, password []byte) {
// 	err := bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
// 	if err != nil {
// 		fmt.Println("Password does not match.")
// 	} else {
// 		fmt.Println("Password matches!")
// 	}
// }

func gohashPass(c echo.Context) error {
	p := new(Pass)
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
