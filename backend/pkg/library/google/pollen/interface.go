package pollen

import "context"

// PollenClient は Google Pollen API クライアントのインターフェース
type PollenClient interface {
	GetForecast(ctx context.Context, req *ForecastRequest) (*ForecastResponse, error)
}

// ForecastRequest は Google Pollen API へのリクエストパラメータ
type ForecastRequest struct {
	Lat  float64
	Lng  float64
	Days int
}

// ForecastResponse は Google Pollen API からのレスポンス
type ForecastResponse struct {
	DailyForecasts []DailyForecast
}

// DailyForecast は1日分の花粉情報
type DailyForecast struct {
	Plants []Plant
}

// Plant は花粉種ごとの指数情報
type Plant struct {
	Code     string
	InSeason bool
	Level    int
}
