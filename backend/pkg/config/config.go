package config

import (
	"context"
	"fmt"
	"os"

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

// TursoDBConfig は LoadSecrets で設定される
type TursoDBConfig struct {
	URL       string `json:"database_url"`
	AuthToken string `json:"auth_token"`
}

// GoogleConfig は LoadSecrets で設定される
type GoogleConfig struct {
	PollenAPIKey string
}

func Load() (*Config, error) {
	cfg, err := env.ParseAs[Config]()
	if err != nil {
		return nil, fmt.Errorf("parse config: %w", err)
	}

	if !cfg.App.Env.IsValid() {
		return nil, fmt.Errorf("invalid APP_ENV: %q (must be one of: %s, %s)", cfg.App.Env, EnvDev, EnvPrd)
	}

	return &cfg, nil
}

// LoadSecrets は環境に応じて機密値を取得する。
// 本番環境では Secrets Manager から取得し、開発環境では環境変数から直接読む。
func (c *Config) LoadSecrets(ctx context.Context) error {
	if c.App.Env.IsProduction() {
		return c.loadFromSecretsManager(ctx)
	}
	return c.loadFromEnv()
}

func (c *Config) loadFromEnv() error {
	c.TursoDB.URL = os.Getenv("TURSO_DATABASE_URL")
	c.TursoDB.AuthToken = os.Getenv("TURSO_AUTH_TOKEN")
	c.Google.PollenAPIKey = os.Getenv("GOOGLE_POLLEN_API_KEY")
	return nil
}
