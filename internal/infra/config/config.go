package config

import "os"

type Config struct {
	Env      string
	HTTPPort string
}

func Load() *Config {
	cfg := &Config{
		Env:      getEnv("SERVICE_ENV", "dev"),
		HTTPPort: getEnv("HTTP_PORT", "8080"),
	}
	return cfg
}

func getEnv(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return fallback
}
