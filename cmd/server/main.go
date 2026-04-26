package main

import (
	"gininitial/internal/api"
	config "gininitial/internal/config"
	"gininitial/internal/database"
	"strconv"

	"log/slog"
	"os"
)

var logger *slog.Logger = config.InitLogger("startup-service")

// @title           Gin Initial Blueprint API
// @version         1.0
// @description     This is a blueprint server for a Gin API structure.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8081
// @BasePath  /

/*
migrate-only mode for deploy pipelines / Kubernetes init containers:
./server migrate
Main process runs migrations then starts HTTP unless SKIP_DB_MIGRATE=true (when migrate already ran).
*/
func main() {
	if len(os.Args) > 1 && os.Args[1] == "migrate" {
		database.RunMigrateCommand()
		return
	}

	port := strconv.Itoa(config.ENV.App.APP_PORT)
	if port == "" {
		logger.Error("PORT is not set")
		os.Exit(1)
	}

	db := database.InitDB()
	defer func() {
		if err := db.Close(); err != nil {
			logger.Error("Database close failed", slog.String("error", err.Error()))
		}
	}()

	if err := database.MigrateIfEnabled(db); err != nil {
		logger.Error("Migration failed", slog.String("error", err.Error()))
		os.Exit(1)
	}

	apiLogger := config.InitLogger("Api")

	deps := api.RouterDependencies{
		Logger: apiLogger,
		DB:     db,
	}

	healthPort := strconv.Itoa(config.ENV.App.APP_HEALTH_CHECK_PORT)
	if healthPort == "" {
		logger.Error("HEALTH_CHECK_PORT is not set")
		os.Exit(1)
	}
	if healthPort == port {
		logger.Error("PORT and HEALTH_CHECK_PORT must differ", slog.String("port", port), slog.String("healthCheckPort", healthPort))
		os.Exit(1)
	}

	apiEngine := api.SetupRouter(deps)
	healthEngine := api.SetupHealthRouter(deps)

	go func() {
		logger.Info("Starting health server", slog.String("addr", ":"+healthPort))
		if err := healthEngine.Run(":" + healthPort); err != nil {
			logger.Error("Health server failed", slog.String("error", err.Error()))
			os.Exit(1)
		}
	}()

	logger.Info("Starting API server", slog.String("addr", ":"+port))
	if err := apiEngine.Run(":" + port); err != nil {
		logger.Error("API server failed", slog.String("error", err.Error()))
		os.Exit(1)
	}
}
