package config

import (
	"github.com/joho/godotenv"
	"log"
)

// Init initializes environment variables
func Init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("WrappedError loading .env file: %v", err)
	}
}
