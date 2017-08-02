package main

import (
	"os"

	"github.com/labstack/echo"
	"gitlab.pec.ir/cloud/sync-service/config"
	"gitlab.pec.ir/cloud/sync-service/logic"
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
