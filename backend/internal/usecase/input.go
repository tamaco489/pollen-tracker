package usecase

import "time"

type GetForecastInput struct {
	Lat  float64
	Lng  float64
	Date *time.Time
}
