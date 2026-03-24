package main

import (
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// setupLogger initializes a structured JSON logger
func setupLogger() *slog.Logger {
	// Read log level from env, default to Info
	logLevel := slog.LevelInfo
	if envLevel := os.Getenv("LOG_LEVEL"); envLevel != "" {
		if err := logLevel.UnmarshalText([]byte(envLevel)); err != nil {
			// fallback to Info if invalid
			logLevel = slog.LevelInfo
		}
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: logLevel,
	}))

	// Set as default logger for standard library compatibility
	slog.SetDefault(logger)
	return logger
}

// slogMiddleware logs a Gin HTTP request in JSON format using slog
func slogMiddleware(logger *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// Process request
		c.Next()

		// Gather information after processing
		latency := time.Since(start)
		status := c.Writer.Status()

		msg := "HTTP Request"
		logAttrs := []slog.Attr{
			slog.Int("status", status),
			slog.String("method", c.Request.Method),
			slog.String("path", c.Request.URL.Path),
			slog.String("ip", c.ClientIP()),
			slog.Duration("latency", latency),
			slog.String("user_agent", c.Request.UserAgent()),
		}

		if len(c.Errors) > 0 {
			logAttrs = append(logAttrs, slog.String("errors", c.Errors.String()))
			logger.LogAttrs(c.Request.Context(), slog.LevelError, msg, logAttrs...)
		} else if status >= 400 {
			logger.LogAttrs(c.Request.Context(), slog.LevelWarn, msg, logAttrs...)
		} else {
			logger.LogAttrs(c.Request.Context(), slog.LevelInfo, msg, logAttrs...)
		}
	}
}

func setupRouter(logger *slog.Logger) *gin.Engine {
	// Use New() instead of Default() to avoid the default text logger
	r := gin.New()

	// Add recovery middleware to recover from panics and log them
	r.Use(gin.Recovery())

	// Add our custom structured JSON logger middleware
	r.Use(slogMiddleware(logger))

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		logger.Debug("Ping endpoint hit", slog.String("endpoint", "/ping"))
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	return r
}

func main() {
	// Load environment variables from .env file first
	err := godotenv.Load()

	// Initialize structured logger
	logger := setupLogger()

	if err != nil {
		logger.Warn("No .env file found or error loading it. Using environment variables.")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := setupRouter(logger)

	logger.Info("Starting server", slog.String("port", port))
	if err := r.Run(":" + port); err != nil {
		logger.Error("Server failed to start", slog.String("error", err.Error()))
		os.Exit(1)
	}
}
