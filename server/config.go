package server

import (
	"fmt"
	"log"

	env "github.com/Netflix/go-env"
	"github.com/joho/godotenv"
)

type config struct {
	PostgresURL string `env:"POSTGRES_URL,required=true"`
}

func loadConfig(localMode bool) (*config, error) {
	if localMode {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	cfg := new(config)
	if _, err := env.UnmarshalFromEnviron(cfg); err != nil {
		return nil, fmt.Errorf("failed to parse environment variables: %w", err)
	}
	fmt.Println("Loaded environment successfully.")
	return cfg, nil
}
