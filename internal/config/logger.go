package config

import (
	"log/slog"
	"os"
	"strings"
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

// InitLogger returns the shared JSON logger, optionally with a "service" attribute.
// Call InitLogger() with no args, InitLogger(""), or omit the name to skip the attribute.
func InitLogger(service ...string) *slog.Logger {
	base := setupLogger()
	if len(service) == 0 {
		return base
	}
	name := strings.TrimSpace(service[0])
	if name == "" {
		return base
	}
	return base.With(slog.String("service", name))
}
