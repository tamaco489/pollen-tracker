package usecase

import (
	"context"
	"fmt"

	"github.com/oapi-codegen/runtime/types"

	"github.com/tamaco489/pollen-tracker/backend/internal/gen"
	"github.com/tamaco489/pollen-tracker/backend/internal/infrastructure/datastore"
	"github.com/tamaco489/pollen-tracker/backend/pkg/errors/sentinel"
	"github.com/tamaco489/pollen-tracker/backend/pkg/utils"
)

type getStatsUseCase struct {
	repo datastore.GetStatsRepository
}

func NewGetStats(repo datastore.GetStatsRepository) GetStatsUseCase {
	return &getStatsUseCase{repo: repo}
}

func (uc *getStatsUseCase) GetStats(ctx context.Context, input gen.GetStatsParams) (*gen.StatsResponse, error) {
	if !input.Period.Valid() {
		return nil, fmt.Errorf("%w: period must be weekly or monthly", sentinel.ErrInvalidInput)
	}

	from, to := utils.DefaultDateRange()
	if input.From != nil {
		from = input.From.Time
	}
	if input.To != nil {
		to = input.To.Time
	}

	var (
		items []gen.StatsItem
		err   error
	)

	switch input.Period {
	case gen.Weekly:
		rows, e := uc.repo.GetWeeklyStats(ctx, from, to)
		if e != nil {
			return nil, fmt.Errorf("get weekly stats: %w", e)
		}
		items = make([]gen.StatsItem, 0, len(rows))
		for _, r := range rows {
			items = append(items, gen.StatsItem{
				StartDate:      types.Date{Time: r.StartDate},
				EndDate:        types.Date{Time: r.EndDate},
				AvgSneezing:    float32(r.AvgSneezing),
				AvgRunny:       float32(r.AvgRunny),
				AvgItchy:       float32(r.AvgItchy),
				AvgPollenLevel: float32(r.AvgPollenLevel),
				Count:          int32(r.Count),
			})
		}
	case gen.Monthly:
		rows, e := uc.repo.GetMonthlyStats(ctx, from, to)
		if e != nil {
			return nil, fmt.Errorf("get monthly stats: %w", e)
		}
		items = make([]gen.StatsItem, 0, len(rows))
		for _, r := range rows {
			items = append(items, gen.StatsItem{
				StartDate:      types.Date{Time: r.StartDate},
				EndDate:        types.Date{Time: r.EndDate},
				AvgSneezing:    float32(r.AvgSneezing),
				AvgRunny:       float32(r.AvgRunny),
				AvgItchy:       float32(r.AvgItchy),
				AvgPollenLevel: float32(r.AvgPollenLevel),
				Count:          int32(r.Count),
			})
		}
	}

	if err != nil {
		return nil, err
	}

	return &gen.StatsResponse{
		Period: input.Period,
		Items:  items,
	}, nil
}
