package configs

import (
	"go-adv/3-validation-api/pkg/env"

	"github.com/joho/godotenv"
)

type Config struct {
	Email      string
	Password   string
	Port       string
	Protocol   string
	Domain     string
	JsonDbPath string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}

	return &Config{
		Email:      env.Getenv("EMAIL", ""),
		Password:   env.Getenv("PASSWORD", ""),
		Port:       env.Getenv("PORT", "8081"),
		Protocol:   env.Getenv("PROTOCOL", "http://"),
		Domain:     env.Getenv("DOMAIN", "localhost"),
		JsonDbPath: env.Getenv("JSON_DB_PATH", ""),
	}
}
