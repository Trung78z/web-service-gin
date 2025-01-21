package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBConn      string
	AppPort     string
	RedisURL    string
	Development bool
}

func Load() *Config {
	// Load file .env
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found. Using environment variables or defaults.")
	}

	return &Config{
		DBConn:      getEnv("DATABASE_URL", "default"),
		AppPort:     getEnv("PORT", "8080"),
		RedisURL:    getEnv("REDIS_URL", "default"),
		Development: getEnv("DEVELOPMENT", "true") == "true",
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
