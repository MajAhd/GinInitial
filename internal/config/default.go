package config

func Default() *Schema {
	return &Schema{
		App: AppConfig{
			APP_NAME:              "GinApplication",
			APP_HOST:              "localhost",
			APP_PORT:              8080,
			APP_HEALTH_CHECK_PORT: 8040,
			APP_LOGLEVEL:          "INFO",
			SKIP_DB_MIGRATE:       false,
			GIN_MODE:              "debug",
		},
		DB: DatabaseConfig{},
	}
}
