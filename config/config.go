// config/config.go
package config

import (
	"log"
	"os"

	"github.com/Shrihari8726/Gonic-Gin-REST-API-2024/godotenv"
)

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func GetEnv(key string) string {
	return os.Getenv(key)
}

func GetMongoURI() string {
	return os.Getenv("MONGO_URI")
}

func GetJWTSecret() string {
	return os.Getenv("JWT_SECRET")
}
