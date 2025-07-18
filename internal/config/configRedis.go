package taskifyConfig

import taskifyEnv "github.com/raihandotmd/taskify/pkg/env"

type RedisConfig struct {
	Addr     string `envconfig:"REDIS_ADDR" required:"true"`
	Password string `envconfig:"REDIS_PASSWORD" default:""`
	DB       int    `envconfig:"REDIS_DB" default:"0"`
}

func NewRedisConfig() (RedisConfig, error) {
	config := RedisConfig{}
	if err := taskifyEnv.New(&config); err != nil {
		return RedisConfig{}, err
	}
	return config, nil
}
