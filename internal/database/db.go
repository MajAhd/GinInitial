package database

import (
	"context"
	"database/sql"
	config "gininitial/internal/config"
	"log/slog"
	"net/url"
	"os"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
)

var logger *slog.Logger = config.InitLogger("database-service")

func postgresDSNFromEnv() string {
	host := config.ENV.DB.DB_HOSTNAME
	port := config.ENV.DB.DB_PORT
	user := config.ENV.DB.DB_USERNAME
	pass := config.ENV.DB.DB_PASSWORD
	dbname := config.ENV.DB.DB_DATABASE
	sslmode := config.ENV.DB.DB_SSL_DISABLED
	if sslmode == "true" {
		sslmode = "disable"
	}
	u := &url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(user, pass),
		Host:   host + ":" + port,
		Path:   "/" + dbname,
	}
	q := u.Query()
	q.Set("sslmode", sslmode)
	u.RawQuery = q.Encode()
	return u.String()
}

// InitDB initializes PostgreSQL connection and configures Bun ORM
func InitDB() *bun.DB {
	dsn := postgresDSNFromEnv()
	if dsn == "" {
		logger.Error("Database connection failed", slog.String("error", "DATABASE_URL is not set"))
		os.Exit(1)
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
