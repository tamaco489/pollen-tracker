package utils

import "time"

func DefaultDateRange() (from, to time.Time) {
	now := time.Now().UTC().Truncate(24 * time.Hour)
	return now.AddDate(0, 0, -30), now
}
