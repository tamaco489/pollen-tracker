package dto

import (
	"time"

	"github.com/google/uuid"
)

type GetSymptoms struct {
	ID             uuid.UUID
	Date           time.Time
	Sneezing       int32
	Runny          int32
	Itchy          int32
	PollenLevel    int32
	TookMedication bool
	Note           string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type CreateSymptoms struct {
	ID             uuid.UUID
	Date           time.Time
	Sneezing       int32
	Runny          int32
	Itchy          int32
	PollenLevel    int32
	TookMedication bool
	Note           string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type UpdateSymptoms struct {
	ID             uuid.UUID
	Sneezing       int32
	Runny          int32
	Itchy          int32
	PollenLevel    int32
	TookMedication bool
	Note           string
	UpdatedAt      time.Time
}

type StatsItem struct {
	StartDate      time.Time
	EndDate        time.Time
	AvgSneezing    float64
	AvgRunny       float64
	AvgItchy       float64
	AvgPollenLevel float64
	Count          int64
}
