package main

import (
	"gininitial/internal/api"
	"gininitial/internal/database"
	"gininitial/internal/models"

	"github.com/joho/godotenv"

	"log/slog"
	"os"
)

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

func setupLogger() *slog.Logger {
	logLevel := slog.LevelInfo
	if envLevel := os.Getenv("LOG_LEVEL"); envLevel != "" {
		if err := logLevel.UnmarshalText([]byte(envLevel)); err != nil {
			logLevel = slog.LevelInfo
		}
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: logLevel,
	}))

	slog.SetDefault(logger)
	return logger
}

func main() {
	err := godotenv.Load()

	logger := setupLogger()

	if err != nil {
		logger.Warn("No .env file found or error loading it. Using environment variables.")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	db := database.InitDB(logger)
	defer db.Close()

	// Automatically run our DB testing schema migrations
	if err := database.Migrate(db, (*models.User)(nil)); err != nil {
		logger.Error("Migration Error", slog.String("error", err.Error()))
	}

	deps := api.RouterDependencies{
		Logger: logger,
		DB:     db,
	}

	r := api.SetupRouter(deps)

	logger.Info("Starting server", slog.String("port", port))
	if err := r.Run(":" + port); err != nil {
		logger.Error("Server failed to start", slog.String("error", err.Error()))
		os.Exit(1)
	}
}
