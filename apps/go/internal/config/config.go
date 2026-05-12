package config

import "os"

type Config struct {
	Host string
	Port string
}

func Load() Config {
	return Config{
		Host: getenv("HOST", "0.0.0.0"),
		Port: getenv("PORT", "8080"),
	}
}

func getenv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}

	return fallback
}
