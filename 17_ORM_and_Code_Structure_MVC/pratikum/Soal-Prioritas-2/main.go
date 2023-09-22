package main

import (
	"soal-prioritas-2/config"
	"soal-prioritas-2/routes"
)

func main() {

	config.InitDB()
	e := routes.New()
	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))

}
