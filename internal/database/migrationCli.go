package database

import (
	"gininitial/internal/models"
	"log/slog"
	"os"

	"github.com/uptrace/bun"
)

func RunMigrateCommand() {
	db := InitDB()
	defer func() {
		if err := db.Close(); err != nil {
			logger.Error("Database close failed", slog.String("error", err.Error()))
		}
	}()
	if err := Migrate(db, (*models.User)(nil)); err != nil {
		logger.Error("Migration failed", slog.String("error", err.Error()))
		os.Exit(1)
	}
	logger.Info("Migrations completed successfully")
}

func MigrateIfEnabled(db *bun.DB) error {
	if os.Getenv("SKIP_DB_MIGRATE") == "true" {
		logger.Info("SKIP_DB_MIGRATE is set: skipping database migration")
		return nil
	}
	return Migrate(db, (*models.User)(nil))
}
