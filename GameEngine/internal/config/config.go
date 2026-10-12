package config

import "os"

type Config struct {
	AppAddr     string
	DatabaseURL string
}

func Load() Config {
	return Config{
		AppAddr: getEnv(
			"APP_ADDR",
			":8080",
		),

		DatabaseURL: getEnv(
			"DATABASE_URL",
			"postgres://dnd:dnd@localhost:5432/dnd?sslmode=disable",
		),
	}

}

func getEnv(
	key string,
	fallback string,
) string {
	value := os.Getenv(key)

	if value == "" {
		return fallback
	}

	return value
}
