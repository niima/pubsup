package main

import (
	"os"

	"git.raad.cloud/cloud/sync-service/config"
	"git.raad.cloud/cloud/sync-service/logic"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	config.RegisterMiddlewares(e)
	config.RegisterRoutes(e)

	//set default port or load frop env
	port := "9002"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}

	go logic.ConsumerPool()
	go e.Logger.Fatal(e.Start(":" + port))

}
