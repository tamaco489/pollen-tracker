package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"
)

type AuthorizerConfig struct {
	App       AppConfig
	SecretARN string `env:"SECRET_ARN,required,notEmpty"`
}

func LoadAuthorizer() (*AuthorizerConfig, error) {
	cfg, err := env.ParseAs[AuthorizerConfig]()
	if err != nil {
		return nil, fmt.Errorf("parse config: %w", err)
	}

	if !cfg.App.Env.IsValid() {
		return nil, fmt.Errorf("invalid APP_ENV: %q (must be one of: %s, %s)", cfg.App.Env, EnvDev, EnvPrd)
	}

	return &cfg, nil
}
