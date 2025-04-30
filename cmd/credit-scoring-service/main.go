package main

import (
	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"

	"github.com/bahtey101/credit-scoring-service/internal/app"
	"github.com/bahtey101/credit-scoring-service/internal/config"
	"github.com/bahtey101/credit-scoring-service/pkg/logging"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		logrus.Fatal("Error loading .env file: ", env.ErrNotAStructPtr)
	}

	cfg := new(config.Config)

	if err := env.Parse(cfg); err != nil {
		logrus.Fatal("Failed to retreive env variables: ", err)
	}

	logging.SetLogging(cfg.LogLevel)

	if err := app.Run(cfg); err != nil {
		logrus.Fatal("Error running servicde: ", err)
	}

}
