package controllers

import (
	"net/http"

	"github.com/labstack/echo"
)

type person struct {
	name  string
	Email string
}

//EchoHello is test
func EchoHello(c echo.Context) error {
	// return c.JSON(http.StatusOK, "ok")
	p := person{name: "hi", Email: "go"}
	return c.JSON(http.StatusOK, p)
}

//HelloName is hello
func HelloName(c echo.Context) error {
	name := c.QueryParam("name")
	return c.String(http.StatusOK, "hello "+name)
}
