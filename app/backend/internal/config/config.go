package config

import "os"

type Config struct{ AppPort, AppEnv, CORSOrigin, DatabaseURL string }

func Load() Config {
	return Config{
		AppPort:     getEnv("APP_PORT", "8080"),
		AppEnv:      getEnv("APP_ENV", "development"),
		CORSOrigin:  getEnv("CORS_ORIGIN", "http://localhost:5173"),
		DatabaseURL: os.Getenv("DATABASE_URL"),
	}
}

func (c Config) DatabaseDSN() string {
	if c.DatabaseURL != "" {
		return c.DatabaseURL
	}
	return "postgres://" + getEnv("POSTGRES_USER", "cloud_platform") + ":" + getEnv("POSTGRES_PASSWORD", "cloud_platform_dev") + "@" + getEnv("POSTGRES_HOST", "localhost") + ":" + getEnv("POSTGRES_PORT", "5432") + "/" + getEnv("POSTGRES_DB", "cloud_platform")
}
func getEnv(k, fallback string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return fallback
}
