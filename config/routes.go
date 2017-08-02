package config

import (
	"github.com/labstack/echo"
	"gitlab.pec.ir/cloud/sync-service/controllers"
)

//RegisterRoutes register routes
func RegisterRoutes(e *echo.Echo) {
	e.GET("/", controllers.EchoHello)
	e.GET("/hello", controllers.HelloName)
	e.POST("/hello", controllers.HelloPost)
}
