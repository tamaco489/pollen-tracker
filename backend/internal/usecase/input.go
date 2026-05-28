package usecase

import "time"

type GetForecastInput struct {
	Lat  float64
	Lng  float64
	Date *time.Time
}

type PostSymptomsInput struct {
	Date           time.Time
	Sneezing       int32
	Runny          int32
	Itchy          int32
	PollenLevel    int32
	TookMedication bool
	Note           string
}

type GetSymptomsInput struct {
	From time.Time
	To   time.Time
}
