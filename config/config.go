package config

import "os"

type DBconfig struct {
	DB string
}

func CreateConfig() *DBconfig {
	return &DBconfig{
		DB: getEnv("DB", ""),
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
