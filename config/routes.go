package config

import (
	"github.com/labstack/echo"
	"gitlab.pec.ir/cloud/sync-service/controllers"
)

//RegisterRoutes register routes
func RegisterRoutes(e *echo.Echo) {

	//publish
	e.POST("/publish", controllers.Publish)
}
