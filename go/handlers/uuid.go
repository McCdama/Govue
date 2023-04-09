package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	uuid "github.com/satori/go.uuid"
)

type IDs struct {
	Index int       `json:"index" xml:"index" form:"index" query:"index"`
	Uuid  uuid.UUID `json:"uuid" xml:"uuid" form:"uuid" query:"uuid"`
}

func Gouuid(c echo.Context) error {

	numberRequested, _ := strconv.Atoi(c.FormValue("numberRequested"))

	ids := make([]IDs, numberRequested)

	for i := 0; i < numberRequested; i++ {
		generatedUUID := uuid.NewV4()
		ids[i] = IDs{
			Index: i,
			Uuid:  generatedUUID,
		}
	}

	return c.JSON(http.StatusOK, ids)

}
