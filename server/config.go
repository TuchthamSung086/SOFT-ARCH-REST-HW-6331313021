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
	fmt.Println("trying to print env1")
	if _, err := env.UnmarshalFromEnviron(cfg); err != nil {
		return nil, fmt.Errorf("failed to parse environment variables: %w", err)
	}
	fmt.Println("trying to print env2")
	fmt.Println(cfg.PostgresURL)
	return cfg, nil
}
