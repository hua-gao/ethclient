package ethutil

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnvParam(name string) (value string) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv(name)
}
