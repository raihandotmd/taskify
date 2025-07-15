package taskifyConfig

import taskifyEnv "github.com/raihandotmd/taskify/pkg/env"

type PGConfig struct {
	Host     string `envconfig:"GORM_HOST" required:"true"`
	Port     string `envconfig:"GORM_PORT" required:"true"`
	User     string `envconfig:"GORM_USER" required:"true"`
	Password string `envconfig:"GORM_PASSWORD" required:"true"`
	DBName   string `envconfig:"GORM_DB" required:"true"`
}

func NewPGConfig() (PGConfig, error) {
	config := PGConfig{}
	if err := taskifyEnv.New(&config); err != nil {
		return PGConfig{}, err
	}
	return config, nil
}
