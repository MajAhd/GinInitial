package config

type Schema struct {
	App AppConfig      `envPrefix:""`
	DB  DatabaseConfig `envPrefix:""`
}

type AppConfig struct {
	APP_NAME              string `env:"APP_NAME,required"`
	APP_HOST              string `env:"APP_HOST"`
	APP_PORT              int    `env:"APP_PORT,required"`
	APP_HEALTH_CHECK_PORT int    `env:"APP_HEALTH_CHECK_PORT,required"`
	APP_LOGLEVEL          string `env:"APP_LOGLEVEL,required"`
	SKIP_DB_MIGRATE       bool   `env:"SKIP_DB_MIGRATE"`
	GIN_MODE              string `env:"GIN_MODE"`
}

type DatabaseConfig struct {
	DB_HOSTNAME     string `env:"DB_HOSTNAME,required"`
	DB_PORT         string `env:"DB_PORT,required"`
	DB_USERNAME     string `env:"DB_USERNAME,required"`
	DB_PASSWORD     string `env:"DB_PASSWORD,required"`
	DB_DATABASE     string `env:"DB_DATABASE,required"`
	DB_SSL_DISABLED string `env:"DB_SSL_DISABLED,required"`
}
