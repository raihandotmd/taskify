package taskifyConfig

import (
	"time"

	taskifyEnv "github.com/raihandotmd/taskify/pkg/env"
)

type TokenConfig struct {
	TokenSecret     string        `envconfig:"TOKEN_SECRET" required:"true"`
	TokenExpiration time.Duration `envconfig:"TOKEN_EXPIRATION" required:"true"`
}

func NewTokenConfig() (TokenConfig, error) {
	config := TokenConfig{}
	if err := taskifyEnv.New(&config); err != nil {
		return TokenConfig{}, err
	}
	return config, nil
}
