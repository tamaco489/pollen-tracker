package main

import (
	"context"
	"crypto/subtle"
	"fmt"
	"os"
	"sync"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

var (
	mu        sync.RWMutex
	cachedKey []byte
)

// loadAPIKey は Secrets Manager から API キーを取得してメモリにキャッシュする。
// 初回コールドスタート時のみ Secrets Manager を呼び出し、以降はキャッシュを返す。
// 取得失敗時はキャッシュせず、次回呼び出し時に再試行する。
func loadAPIKey(ctx context.Context) ([]byte, error) {
	mu.RLock()
	if len(cachedKey) > 0 {
		key := cachedKey
		mu.RUnlock()
		return key, nil
	}
	mu.RUnlock()

	mu.Lock()
	defer mu.Unlock()

	if len(cachedKey) > 0 {
		return cachedKey, nil
	}

	secretARN := os.Getenv("SECRET_ARN")
	if secretARN == "" {
		return nil, fmt.Errorf("SECRET_ARN is not set")
	}

	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, fmt.Errorf("load AWS config: %w", err)
	}

	client := secretsmanager.NewFromConfig(cfg)
	out, err := client.GetSecretValue(ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretARN),
	})
	if err != nil {
		return nil, fmt.Errorf("get secret value: %w", err)
	}

	if out.SecretString == nil {
		return nil, fmt.Errorf("secret has no string value")
	}

	cachedKey = []byte(*out.SecretString)
	return cachedKey, nil
}

func handler(ctx context.Context, event events.APIGatewayV2CustomAuthorizerV2Request) (events.APIGatewayV2CustomAuthorizerSimpleResponse, error) {
	apiKey, err := loadAPIKey(ctx)
	if err != nil {
		return events.APIGatewayV2CustomAuthorizerSimpleResponse{IsAuthorized: false}, nil
	}

	var requestKey string
	if len(event.IdentitySource) > 0 {
		requestKey = event.IdentitySource[0]
	}

	const equal = 1
	authorized := subtle.ConstantTimeCompare(apiKey, []byte(requestKey)) == equal
	return events.APIGatewayV2CustomAuthorizerSimpleResponse{IsAuthorized: authorized}, nil
}

func main() {
	lambda.Start(handler)
}
