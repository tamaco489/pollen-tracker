package pollen

import (
	"net/http"
	"time"
)

// pollenClient は Google Pollen API への HTTP クライアント実装
type pollenClient struct {
	httpClient *http.Client
	apiKey     string
}

// NewPollenClient は新しい PollenClient を返す
func NewPollenClient(apiKey string) PollenClient {
	return &pollenClient{
		httpClient: &http.Client{Timeout: 10 * time.Second},
		apiKey:     apiKey,
	}
}
