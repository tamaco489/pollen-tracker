package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	App     AppConfig
	TursoDB TursoDBConfig
	Google  GoogleConfig
}

type AppConfig struct {
	Env     Environment `env:"APP_ENV"     envDefault:"dev"`
	Port    string      `env:"APP_PORT"    envDefault:"8080"`
	Project string      `env:"APP_PROJECT" envDefault:"pollen-tracker"`
}

func (a AppConfig) ServiceName() string {
	return fmt.Sprintf("%s-api", a.Project)
}

type TursoDBConfig struct {
	URL       string `env:"TURSO_DATABASE_URL,required,notEmpty"`
	AuthToken string `env:"TURSO_AUTH_TOKEN"`
}

type GoogleConfig struct {
	PollenAPIKey string `env:"GOOGLE_POLLEN_API_KEY"`
}

func Load() (*Config, error) {
	cfg, err := env.ParseAs[Config]()
	if err != nil {
		return nil, fmt.Errorf("parse config: %w", err)
	}

	if !cfg.App.Env.IsValid() {
		return nil, fmt.Errorf("invalid APP_ENV: %q (must be one of: %s, %s)", cfg.App.Env, EnvDev, EnvPrd)
	}

	// 本番環境では認証トークンが必須。
	if cfg.App.Env.IsProduction() && cfg.TursoDB.AuthToken == "" {
		return nil, fmt.Errorf("TURSO_AUTH_TOKEN is required in production")
	}

	// 本番環境では Google Pollen API キーが必須。
	if cfg.App.Env.IsProduction() && cfg.Google.PollenAPIKey == "" {
		return nil, fmt.Errorf("GOOGLE_POLLEN_API_KEY is required in production")
	}

	return &cfg, nil
}
