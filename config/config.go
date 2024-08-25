package config

import (
	"github.com/joho/godotenv"
	"log"
)

// Init initializes environment variables
func Init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Err loading .env file: %v", err)
	}
}
