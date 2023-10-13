package main

import (
	"golang-ai/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	// init routing
	routes.InitRouter(e)
	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
}
