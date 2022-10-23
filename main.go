package main

import (
	"github.com/pilseong/banking/app"
	"github.com/pilseong/banking/logger"
)

func main() {
	logger.Info("Starting the application...")
	app.Start()
}
