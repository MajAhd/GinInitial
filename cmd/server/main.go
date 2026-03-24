package main

import (
	"gininitial/internal/api"
	"github.com/joho/godotenv"

	"log/slog"
	"os"
)

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

	r := api.SetupRouter(logger)

	logger.Info("Starting server", slog.String("port", port))
	if err := r.Run(":" + port); err != nil {
		logger.Error("Server failed to start", slog.String("error", err.Error()))
		os.Exit(1)
	}
}
