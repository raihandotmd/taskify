package taskifyEnv

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

const envFile = ".env"

func New(target any) error {
	if err := godotenv.Load(envFile); err != nil {
		return fmt.Errorf("failed to load .env file: %s", err)
	}

	if err := envconfig.Process("", target); err != nil {
		return fmt.Errorf("failed to process environment variables: %s", err)
	}

	return nil
}
