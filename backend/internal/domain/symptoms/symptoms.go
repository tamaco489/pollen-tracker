package symptoms

import (
	"time"

	"github.com/google/uuid"
)

type Symptom struct {
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
