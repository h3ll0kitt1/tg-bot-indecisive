package config

import (
	"errors"
	"os"
)

type Config struct {
	Token  string
	DbHost string
	DbPort string
}

func New() (Config, error) {
	dbHost := "localhost"
	dbPort := "6379"
	if token := os.Getenv("TELEGRAM_APITOKEN"); token != "" {
		return Config{
			Token:  token,
			DbHost: dbHost,
			DbPort: dbPort,
		}, nil
	}
	return Config{}, errors.New("No telegram token in ENV variable")
}
