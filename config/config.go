package config

import "os"

type Config struct {
	DB string
}

func CreateConfig() *Config {
	return &Config{
		DB: getEnv("DB", ""),
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
