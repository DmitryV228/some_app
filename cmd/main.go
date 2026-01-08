package main

import (
	"log"
	"os"

	"go.uber.org/zap"
)

func main() { // надо прокидывать логгер во все осталье части и создавать сервис / репо и тд
	logger := &zap.Logger{}
	defer logger.Sync()

	var err error
	if os.Getenv("LOGGER_PROD") == "1" {
		logger, err = zap.NewProduction()
		if err != nil {
			log.Fatal("fail start zap logger InfoLevel")
		}
	} else {
		logger, err = zap.NewDevelopment()
		if err != nil {
			log.Fatal("fail start zap logger DebugLevel")
		}
	}

}
