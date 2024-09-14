// config/config.go
package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadConfig() {
	fmt.Println("In Config call")
	// var myEnv map[string]string
	// myEnv, err := godotenv.Read()
	err := godotenv.Load("godotenv.env")
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
