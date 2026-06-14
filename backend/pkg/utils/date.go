package utils

import "time"

// DefaultDateRange は from を 30 日前の 00:00:00 UTC、to を当日の 00:00:00 UTC として返す
func DefaultDateRange() (from, to time.Time) {
	now := time.Now().UTC().Truncate(24 * time.Hour)
	return now.AddDate(0, 0, -30), now
}
