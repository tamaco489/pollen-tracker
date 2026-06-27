package main

import (
	"context"
	"crypto/subtle"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/tamaco489/pollen-tracker/backend/pkg/utils/secret"
)

func handler(ctx context.Context, event events.APIGatewayV2CustomAuthorizerV2Request) (events.APIGatewayV2CustomAuthorizerSimpleResponse, error) {
	apiKey, err := secret.LoadAPIKey(ctx)
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
