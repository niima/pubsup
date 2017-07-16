package config

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

//RegisterMiddlewares Registers Middlewares
func RegisterMiddlewares(e *echo.Echo) {

	//middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

}
