package config

import (
	"log/slog"
	"os"
)

var logger *slog.Logger = InitLogger("service-configs")

var ENV *Schema

func init() {

	switch os.Getenv("GIN_MODE") {
	case "debug":
		logger.Info("Loading Local configuration")
		ENV = Local()
	case "release":
		logger.Info("Loading Prod configuration")
		ENV = Prod()
	default:
		logger.Info("Loading Default configuration")
		ENV = Local()
	}

	logger.Info("Configuration loaded successfully",
		slog.String("APP_NAME", ENV.App.APP_NAME),
		slog.String("ENV", ENV.App.GIN_MODE),
	)
}
