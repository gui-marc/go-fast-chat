package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerAddress string `json:"-"`
	DatabaseURL   string `json:"-"`
	RedisURL      string `json:"-"`
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}

func Load() *Config {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	// Load the environment variables
	APP_PORT := getEnv("APP_PORT", ":8080")
	DATABASE_URL := getEnv("DATABASE_URL", "")
	REDIS_URL := getEnv("REDIS_URL", "")

	return &Config{
		ServerAddress: APP_PORT,
		DatabaseURL:   DATABASE_URL,
		RedisURL:      REDIS_URL,
	}
}
