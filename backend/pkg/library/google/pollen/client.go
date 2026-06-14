package pollen

import (
	"net/http"
	"time"
)

type pollenClient struct {
	httpClient *http.Client
	apiKey     string
}

func NewPollenClient(apiKey string) PollenService {
	return &pollenClient{
		httpClient: &http.Client{Timeout: 10 * time.Second},
		apiKey:     apiKey,
	}
}
