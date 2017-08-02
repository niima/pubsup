package controllers

import (
	"net/http"

	"github.com/labstack/echo"
	"gitlab.pec.ir/cloud/sync-service/models"
)

type person struct {
	name  string
	Email string
}

//EchoHello is test
func EchoHello(c echo.Context) error {
	// return c.JSON(http.StatusOK, "ok")
	p := models.Envelop{Data: "123", Tag: "user"}
	return c.JSON(http.StatusOK, p)
}

//HelloName is hello
func HelloName(c echo.Context) error {
	name := c.QueryParam("name")
	return c.String(http.StatusOK, "hello "+name)
}
