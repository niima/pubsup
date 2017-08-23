package controllers

import (
	"net/http"

	"github.com/labstack/echo"
	"gitlab.pec.ir/cloud/sync-service/logic"
	"gitlab.pec.ir/cloud/sync-service/models"
)

//Publish is used by publishers, they publish data
//to our nsq logic. then we pass data and process it]
//along then publish to subscribers
func Publish(c echo.Context) error {
	e := new(models.Envelop)
	if err := c.Bind(e); err != nil {
		return err
	}
	if e.Tag == "" || e.Data == "" {
		return c.JSON(http.StatusBadRequest, nil)
	}
	logic.Publish(e.Tag, e.URLParam, e.Data)
	return c.JSON(http.StatusOK, e)
}
