package config

import (
	"os"
)

type Config struct {
	MongoURI string
}

func Load() *Config {
	url := "mongodb://127.0.0.1:27017/?directConnection=true&serverSelectionTimeoutMS=2000&appName=mongosh+2.2.6"
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
