package main

import (
	"latihan/config"
	"latihan/routes"
)

func main() {

	config.InitDB()
	e := routes.New()
	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))

}
