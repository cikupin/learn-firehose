package config

import (
	"log"

	"github.com/caarlos0/env/v7"
	"github.com/joho/godotenv"
)

func LoadEnv() {
	godotenv.Load("config/.env")

	err := env.Parse(&Cfg)
	if err != nil {
		log.Fatal(err.Error())
	}
}
