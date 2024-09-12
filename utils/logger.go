// utils/logger.go
package utils

import (
	"log"
	"os"
)

func init() {
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)
}

func LogError(err error) {
	if err != nil {
		log.Println("ERROR:", err)
	}
}

func LogInfo(message string) {
	log.Println("INFO:", message)
}
