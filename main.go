package main

import (
	"github.com/eneassena10/banking/app"
	"github.com/eneassena10/banking/logger"
)

func main() {
	// log.Println("starting our application...")
	logger.Info("Starting our application")
	app.Start()
}
