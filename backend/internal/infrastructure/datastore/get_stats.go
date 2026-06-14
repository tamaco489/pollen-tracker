package datastore

import (
	"context"
	"fmt"
	"time"

	"github.com/tamaco489/pollen-tracker/backend/internal/dto"
	"github.com/tamaco489/pollen-tracker/backend/internal/infrastructure/datastore/gen"
)

func (r *symptomsRepository) GetWeeklyStats(ctx context.Context, from, to time.Time) ([]dto.StatsItem, error) {
	rows, err := r.queries.GetWeeklyStats(ctx, r.db, gen.GetWeeklyStatsParams{
		From: from.Format(time.DateOnly),
		To:   to.Format(time.DateOnly),
	})
	if err != nil {
		return nil, fmt.Errorf("get weekly stats: %w", err)
	}

	result := make([]dto.StatsItem, 0, len(rows))
	for _, row := range rows {
		item, err := parseStatsRow(fmt.Sprintf("%v", row.StartDate), fmt.Sprintf("%v", row.EndDate), row.AvgSneezing, row.AvgRunny, row.AvgItchy, row.AvgPollenLevel, row.Count)
		if err != nil {
			return nil, err
		}
		result = append(result, item)
	}
	return result, nil
}

func (r *symptomsRepository) GetMonthlyStats(ctx context.Context, from, to time.Time) ([]dto.StatsItem, error) {
	rows, err := r.queries.GetMonthlyStats(ctx, r.db, gen.GetMonthlyStatsParams{
		From: from.Format(time.DateOnly),
		To:   to.Format(time.DateOnly),
	})
	if err != nil {
		return nil, fmt.Errorf("get monthly stats: %w", err)
	}

	result := make([]dto.StatsItem, 0, len(rows))
	for _, row := range rows {
		item, err := parseStatsRow(fmt.Sprintf("%v", row.StartDate), fmt.Sprintf("%v", row.EndDate), row.AvgSneezing, row.AvgRunny, row.AvgItchy, row.AvgPollenLevel, row.Count)
		if err != nil {
			return nil, err
		}
		result = append(result, item)
	}
	return result, nil
}

func parseStatsRow(startDate, endDate string, avgSneezing, avgRunny, avgItchy, avgPollenLevel float64, count int64) (dto.StatsItem, error) {
	start, err := time.Parse(time.DateOnly, startDate)
	if err != nil {
		return dto.StatsItem{}, fmt.Errorf("parse start_date %q: %w", startDate, err)
	}
	end, err := time.Parse(time.DateOnly, endDate)
	if err != nil {
		return dto.StatsItem{}, fmt.Errorf("parse end_date %q: %w", endDate, err)
	}
	return dto.StatsItem{
		StartDate:      start,
		EndDate:        end,
		AvgSneezing:    avgSneezing,
		AvgRunny:       avgRunny,
		AvgItchy:       avgItchy,
		AvgPollenLevel: avgPollenLevel,
		Count:          count,
	}, nil
}
