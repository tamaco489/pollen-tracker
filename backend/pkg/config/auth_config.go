package config

import (
	"context"
	"fmt"
	"sync"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/caarlos0/env/v11"

	"github.com/aws/aws-sdk-go-v2/config"
)

type AuthorizerConfig struct {
	Env     Environment `env:"ENV"     envDefault:"dev"`
	Project string      `env:"PROJECT" envDefault:"pollen-tracker"`
	Service string      `env:"SERVICE" envDefault:"authorizer"`

	cache apiKeyCache
}

type apiKeyCache struct {
	mu        sync.RWMutex
	cachedKey []byte
}

func LoadAuthorizer() (*AuthorizerConfig, error) {
	cfg, err := env.ParseAs[AuthorizerConfig]()
	if err != nil {
		return nil, fmt.Errorf("parse config: %w", err)
	}

	if !cfg.Env.IsValid() {
		return nil, fmt.Errorf("invalid ENV: %q (must be one of: %s, %s)", cfg.Env, EnvDev, EnvPrd)
	}

	return &cfg, nil
}

// LoadAPIKey は Secrets Manager から API キーを取得してメモリにキャッシュする。
//
//   - 初回コールドスタート時のみ Secrets Manager を呼び出す
//   - 以降の呼び出しはキャッシュを返す
//   - 取得失敗時はキャッシュせず、次回呼び出し時に再試行する
func (c *AuthorizerConfig) Load(ctx context.Context) ([]byte, error) {
	c.cache.mu.RLock()
	if len(c.cache.cachedKey) > 0 {
		key := c.cache.cachedKey
		c.cache.mu.RUnlock()
		return key, nil
	}
	c.cache.mu.RUnlock()

	c.cache.mu.Lock()
	defer c.cache.mu.Unlock()

	if len(c.cache.cachedKey) > 0 {
		return c.cache.cachedKey, nil
	}

	val, err := c.fetchAuthorizerAPIKey(ctx)
	if err != nil {
		return nil, fmt.Errorf("load api key: %w", err)
	}

	c.cache.cachedKey = []byte(val)
	return c.cache.cachedKey, nil
}

// FetchAuthorizerAPIKey は Secrets Manager から API キーを取得する
func (c *AuthorizerConfig) fetchAuthorizerAPIKey(ctx context.Context) (string, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return "", fmt.Errorf("load AWS config: %w", err)
	}

	client := secretsmanager.NewFromConfig(cfg)
	out, err := client.GetSecretValue(ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(fmt.Sprintf("%s/pollen-tracker/api-key", c.Env)),
	})
	if err != nil {
		return "", fmt.Errorf("get secret value: %w", err)
	}

	if out.SecretString == nil {
		return "", fmt.Errorf("secret has no string value")
	}

	return *out.SecretString, nil
}
