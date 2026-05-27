package usecase

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/tamaco489/pollen-tracker/backend/pkg/errors/sentinel"

	domain_pollen "github.com/tamaco489/pollen-tracker/backend/internal/domain/pollen"
	google_pollen "github.com/tamaco489/pollen-tracker/backend/pkg/library/google/pollen"
)

type UseCase struct {
	uc GetForecastUseCase
}

func NewGetForecast(uc GetForecastUseCase) *UseCase {
	return &UseCase{uc: uc}
}

// GetForecast は指定座標・日付の花粉予報を取得してドメインエンティティで返す
func (uc *UseCase) GetForecast(ctx context.Context, input GetForecastInput) (*domain_pollen.PollenForecast, error) {
	if input.Lat < -90 || input.Lat > 90 || input.Lng < -180 || input.Lng > 180 {
		return nil, fmt.Errorf("%w: lat/lng out of range", sentinel.ErrInvalidInput)
	}

	today := time.Now().UTC().Truncate(24 * time.Hour)
	var targetDate time.Time
	if input.Date != nil {
		targetDate = input.Date.UTC().Truncate(24 * time.Hour)
	} else {
		targetDate = today
	}

	dayOffset := int(targetDate.Sub(today).Hours() / 24)
	if dayOffset < 0 || dayOffset > 4 {
		return nil, fmt.Errorf("%w: date must be within 0-4 days from today", sentinel.ErrInvalidInput)
	}

	req := &google_pollen.ForecastRequest{
		Lat:  input.Lat,
		Lng:  input.Lng,
		Days: dayOffset + 1,
	}

	resp, err := uc.uc.GetForecast(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("get forecast: %w", err)
	}

	if len(resp.DailyForecasts) <= dayOffset {
		return nil, errors.New("no daily forecast for requested date")
	}

	pollenType, level := dominantPollen(resp.DailyForecasts[dayOffset].Plants)

	return &domain_pollen.PollenForecast{
		Date:       targetDate,
		Level:      level,
		PollenType: pollenType,
		SeasonInfo: domain_pollen.SeasonCalendar[pollenType],
	}, nil
}

// dominantPollen は Plants の中から最も高い指数の花粉種を返す
// 対応外コードは無視し、該当なしの場合は CEDAR / level 1 をデフォルトとして返す
func dominantPollen(plants []google_pollen.Plant) (domain_pollen.PollenType, int) {
	bestType := domain_pollen.PollenTypeCedar
	bestLevel := 1

	for _, p := range plants {
		pt, ok := domain_pollen.PlantCodeToType[p.Code]
		if !ok {
			continue
		}
		if p.Level > bestLevel || (p.Level == bestLevel && p.InSeason) {
			bestType = pt
			bestLevel = p.Level
		}
	}

	return bestType, bestLevel
}
