package main

import (
	"restful-api-testing/config"
	"restful-api-testing/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// load config
	var cfg = config.InitConfig()

	// open mysql connection
	dbMysql := config.InitMyqlConn(cfg)

	// init routing
	routes.InitRouter(e, dbMysql)
	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))

}
