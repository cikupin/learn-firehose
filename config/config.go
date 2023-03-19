package config

import (
	"encoding/json"
	"fmt"
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

	byteCfg, _ := json.Marshal(Cfg)
	fmt.Print("[Loaded config] ---> ")
	fmt.Println(string(byteCfg))
	fmt.Println()
}
