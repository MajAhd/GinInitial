package database

import (
	"context"
	"database/sql"
	"log/slog"
	"os"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
)

// InitDB initializes PostgreSQL connection and configures Bun ORM
func InitDB(logger *slog.Logger) *bun.DB {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "postgres://postgres:postgres@localhost:5432/appdb?sslmode=disable"
	}

	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	db := bun.NewDB(sqldb, pgdialect.New())

	// Print all queries to stdout for debugging if BUNDEBUG=1 is set
	db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
		bundebug.FromEnv("BUNDEBUG"),
	))

	if err := db.Ping(); err != nil {
		logger.Error("Database connection failed", slog.String("error", err.Error()))
	} else {
		logger.Info("Database connected successfully")
	}

	return db
}

// Migrate is an auto-migration helper for the blueprint testing
func Migrate(db *bun.DB, models ...interface{}) error {
	ctx := context.Background()
	for _, model := range models {
		_, err := db.NewCreateTable().Model(model).IfNotExists().Exec(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}
