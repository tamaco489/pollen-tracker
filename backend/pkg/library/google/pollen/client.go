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

// NewPollenClient は新しい PollenService を返す
func NewPollenClient(apiKey string) PollenService {
	return &pollenClient{
		httpClient: &http.Client{Timeout: 10 * time.Second},
		apiKey:     apiKey,
	}
}
