package config

import (
	"os"
)

type Config struct {
	MongoURI string
}

func Load() *Config {
	url := "mongodb://localhost:27017"
	return &Config{
		MongoURI: getEnv("MONGO_URI", url),
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
