package config

import (
	"log/slog"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

var logger *slog.Logger = InitLogger("env-config")

func findModuleDir() string {
	dir, err := os.Getwd()
	if err != nil {
		return ""
	}
	for i := 0; i < 10; i++ {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir
		}
		dir = filepath.Join(dir, "..")
	}
	return ""
}

func LoadEnv() {
	svcDir := findModuleDir()
	if svcDir == "" {
		if err := godotenv.Load(".env"); err != nil {
			logger.Warn("No .env loaded", slog.String("error", err.Error()))
		}
		return
	}
	repoEnv := filepath.Join(svcDir, "..", "..", ".env")
	svcEnv := filepath.Join(svcDir, ".env")
	if err := godotenv.Load(repoEnv, svcEnv); err != nil {
		logger.Warn("Partial env load", slog.String("error", err.Error()))
	}
}
