package taskifyConfig

import taskifyEnv "github.com/raihandotmd/taskify/pkg/env"

type App struct {
	Name string `envconfig:"APP_NAME" default:"Taskify"`
	Env  string `envconfig:"APP_ENV" default:"development"`
	Host string `envconfig:"APP_HOST" default:"localhost"`
	Port int    `envconfig:"APP_PORT" default:"8080"`
	Mode string `envconfig:"APP_MODE" default:"debug"`
}

func NewAppConfig() (App, error) {
	config := App{}
	if err := taskifyEnv.New(&config); err != nil {
		return App{}, err
	}
	return config, nil
}
