package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

const (
	appName = "bombom"
)

const (
	envProd = "prod"
	envDev  = "dev"
)

type Config struct {
	Env     string `split_words:"true" required:"true"`
	ClubTag string `split_words:"true" required:"true"`

	BS `split_words:"true" required:"true"`
}

func New() (conf Config, err error) {
	if err = envconfig.Process(appName, &conf); err != nil {
		return conf, fmt.Errorf("envconfig process: %w", err)
	}

	return conf, nil
}

func (c Config) IsProdEnv() bool {
	return c.Env == envProd
}

func (c Config) IsDevEnv() bool {
	return c.Env == envDev
}
