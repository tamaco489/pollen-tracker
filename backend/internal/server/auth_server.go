package server

import (
	"context"
	"crypto/subtle"

	"github.com/aws/aws-lambda-go/events"

	"github.com/tamaco489/pollen-tracker/backend/pkg/config"
	"github.com/tamaco489/pollen-tracker/backend/pkg/logger"
)

type AuthServer struct {
	cfg *config.AuthorizerConfig
	l   *logger.Logger
}

func NewAuthServer(cfg *config.AuthorizerConfig, l *logger.Logger) *AuthServer {
	return &AuthServer{cfg: cfg, l: l}
}

func (a *AuthServer) Run(ctx context.Context, event events.APIGatewayV2CustomAuthorizerV2Request) (events.APIGatewayV2CustomAuthorizerSimpleResponse, error) {
	apiKey, err := a.cfg.Load(ctx)
	if err != nil {
		a.l.ErrorContext(ctx, "failed to load api key", "error", err)
		return events.APIGatewayV2CustomAuthorizerSimpleResponse{IsAuthorized: false}, nil
	}

	var requestKey string
	if len(event.IdentitySource) > 0 {
		requestKey = event.IdentitySource[0]
	}

	// ConstantTimeCompare は一致で 1、不一致で 0 を返す
	authorized := subtle.ConstantTimeCompare(apiKey, []byte(requestKey)) == 1
	return events.APIGatewayV2CustomAuthorizerSimpleResponse{IsAuthorized: authorized}, nil
}
