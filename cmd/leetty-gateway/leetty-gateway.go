package main

import (
	"leetty-gateway/internal/app"
	"leetty-gateway/internal/config"
	"leetty-gateway/internal/logger"
)

func main() {
	logger.Logger.Info("Leetty-Gateway application starting.")
	var conf, err = config.ParseConfig()
	if err != nil {
		logger.Logger.Error("Failed to read configuration file", err)
		panic(err)
	}
	leetty := app.NewApp(conf)
	leetty.Start()
	defer leetty.Close()
}
