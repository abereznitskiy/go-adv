package configs

import (
	"go-adv/4-order-api/pkg/env"
	"log"

	"github.com/joho/godotenv"
)

type Config struct {
	Db DbConfig
}

type DbConfig struct {
	Dsn string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file, using default config")
	}

	return &Config{
		Db: DbConfig{Dsn: env.Getenv("DSN", "")},
	}
}
