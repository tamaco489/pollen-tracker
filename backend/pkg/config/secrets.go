package config

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

func (c *Config) loadFromSecretsManager(ctx context.Context) error {
	if err := c.loadTursoSecret(ctx); err != nil {
		return err
	}
	if err := c.loadPollenSecret(ctx); err != nil {
		return err
	}

	return nil
}

func (c *Config) loadTursoSecret(ctx context.Context) error {
	val, err := fetchSecretValue(ctx, fmt.Sprintf("%s/pollen-tracker/turso/config", c.App.Env))
	if err != nil {
		return fmt.Errorf("fetch turso secret: %w", err)
	}

	if err := json.Unmarshal([]byte(val), &c.TursoDB); err != nil {
		return fmt.Errorf("unmarshal turso secret: %w", err)
	}

	if c.TursoDB.URL == "" {
		return fmt.Errorf("turso database URL is empty")
	}
	if c.TursoDB.AuthToken == "" {
		return fmt.Errorf("turso auth token is empty")
	}
	return nil
}

func (c *Config) loadPollenSecret(ctx context.Context) error {
	val, err := fetchSecretValue(ctx, fmt.Sprintf("%s/pollen-tracker/google/pollen-api-key", c.App.Env))
	if err != nil {
		return fmt.Errorf("fetch pollen secret: %w", err)
	}

	if val == "" {
		return fmt.Errorf("google pollen API key is empty")
	}

	c.Google.PollenAPIKey = val
	return nil
}

func fetchSecretValue(ctx context.Context, secretID string) (string, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return "", fmt.Errorf("load AWS config: %w", err)
	}

	client := secretsmanager.NewFromConfig(cfg)
	out, err := client.GetSecretValue(ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretID),
	})
	if err != nil {
		return "", fmt.Errorf("get secret value (id=%s): %w", secretID, err)
	}

	if out.SecretString == nil {
		return "", fmt.Errorf("secret has no string value (id=%s)", secretID)
	}

	return *out.SecretString, nil
}
