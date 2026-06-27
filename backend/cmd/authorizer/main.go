package main

import (
	"context"
	"crypto/subtle"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/tamaco489/pollen-tracker/backend/pkg/config"
	"github.com/tamaco489/pollen-tracker/backend/pkg/utils/secret"
)

var secretARN string

func handler(ctx context.Context, event events.APIGatewayV2CustomAuthorizerV2Request) (events.APIGatewayV2CustomAuthorizerSimpleResponse, error) {
	apiKey, err := secret.LoadAPIKey(ctx, secretARN)
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
	cfg, err := config.LoadAuthorizer()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	secretARN = cfg.SecretARN

	lambda.Start(handler)
}
