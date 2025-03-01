package configs

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Email    string
	Password string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	return &Config{
		Email:    os.Getenv("EMAIL"),
		Password: os.Getenv("PASSWORD"),
	}
}
