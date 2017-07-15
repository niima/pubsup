package controllers

import (
	"net/http"

	"github.com/labstack/echo"
)

func echoHello(c echo.Context) error {
	return c.JSON(http.StatusOK, "ok")
}
