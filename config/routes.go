package config

import (
	"git.raad.cloud/cloud/sync-service/controllers"
	"github.com/labstack/echo"
)

//RegisterRoutes register routes
func RegisterRoutes(e *echo.Echo) {

	//publish
	e.POST("/publish", controllers.Publish)
}
