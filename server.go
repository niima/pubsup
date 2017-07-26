package main

// "os"
import (
	"gitlab.pec.ir/cloud/sync-service/logic"
)

// "github.com/labstack/echo"
// "gitlab.pec.ir/cloud/sync-service/config"
//"gitlab.pec.ir/cloud/sync-service/db"

func main() {
	// e := echo.New()
	// config.RegisterMiddlewares(e)
	// config.RegisterRoutes(e)

	// //set default port or load frop env
	// port := "9002"
	// if os.Getenv("PORT") != "" {
	// 	port = os.Getenv("PORT")
	// }
	// e.Logger.Fatal(e.Start(":" + port))

	//------------bolt
	// db.SetData([]byte("hi"), []byte("bye"))
	// data, err := db.GetData([]byte("hi"))
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(data)
	logic.Producer("event", []byte("hi im nima"))
	logic.Consumer()
}
