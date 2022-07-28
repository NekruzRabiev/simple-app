package config

import (
	"github.com/joho/godotenv"
	"os"
)

func parseEnv(envFilePath string) error {
	return godotenv.Load(envFilePath)
}

func setFromEnv(cfg *Config) {
	cfg.Postgres.Username = os.Getenv("PG_USERNAME")
	cfg.Postgres.Password = os.Getenv("PG_PASSWORD")
	cfg.Jwt.Salt = os.Getenv("JWT_SALT")
}