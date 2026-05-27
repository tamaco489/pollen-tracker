package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/tamaco489/pollen-tracker/backend/pkg/errors/sentinel"

	domain_pollen "github.com/tamaco489/pollen-tracker/backend/internal/domain/pollen"
	google_pollen "github.com/tamaco489/pollen-tracker/backend/pkg/library/google/pollen"
)

type getForecastUseCase struct {
	svc google_pollen.PollenService
}

func NewGetForecast(svc google_pollen.PollenService) GetForecastUseCase {
	return &getForecastUseCase{svc: svc}
}

const (
	// Google Pollen API がサポートする座標の範囲
	minLat float64 = -90
	maxLat float64 = 90
	minLng float64 = -180
	maxLng float64 = 180

	// Google Pollen API がサポートする予報日数の範囲 (今日から 0〜4 日後)
	minDayOffset = 0
	maxDayOffset = 4
)

// GetForecast は指定座標・日付の花粉予報を取得してドメインエンティティで返す
func (uc *getForecastUseCase) GetForecast(ctx context.Context, input GetForecastInput) (*domain_pollen.PollenForecast, error) {
	if err := uc.validate(input.Lat, input.Lng); err != nil {
		return nil, err
	}

	// 対象日付を決定する: 未指定なら今日、指定があればその日付を UTC 00:00:00 に揃える
	today := time.Now().UTC().Truncate(24 * time.Hour)
	targetDate := today
	if input.Date != nil {
		targetDate = input.Date.UTC().Truncate(24 * time.Hour)
	}

	// 今日から対象日付までの日数を算出し、API がサポートする範囲 (0〜4 日後) か検証する
	dayOffset := int(targetDate.Sub(today).Hours() / 24)
	if dayOffset < minDayOffset || dayOffset > maxDayOffset {
		return nil, fmt.Errorf(
			"%w: date must be within %d-%d days from today",
			sentinel.ErrInvalidInput,
			minDayOffset,
			maxDayOffset,
		)
	}

	// Days は「今日から何日分取得するか」を指定する: 対象日付の index に到達するために +1 する
	request := &google_pollen.ForecastRequest{
		Lat:  input.Lat,
		Lng:  input.Lng,
		Days: dayOffset + 1,
	}

	// Google Pollen API から予報を取得
	resp, err := uc.svc.GetForecast(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("get forecast: %w", err)
	}

	// 対象日付の予報データが API レスポンスに含まれない場合はエラーを返す
	if len(resp.DailyForecasts) <= dayOffset {
		return nil, fmt.Errorf("%w: no daily forecast for requested date", sentinel.ErrNotFound)
	}

	pollenType, level := uc.dominantPollen(resp.DailyForecasts[dayOffset].Plants)

	return &domain_pollen.PollenForecast{
		Date:       targetDate,
		Level:      level,
		PollenType: pollenType,
		SeasonInfo: domain_pollen.SeasonCalendar[pollenType],
	}, nil
}

func (uc *getForecastUseCase) validate(lat, lng float64) error {
	if lat < minLat || lat > maxLat || lng < minLng || lng > maxLng {
		return fmt.Errorf("%w: lat/lng out of range", sentinel.ErrInvalidInput)
	}
	return nil
}

// dominantPollen は Plants の中から最も高い指数の花粉種を返す
//
// 対応外コードは無視し、該当なしの場合は CEDAR / level 0 をデフォルトとして返す
func (uc *getForecastUseCase) dominantPollen(plants []google_pollen.Plant) (domain_pollen.PollenType, int) {
	bestType := domain_pollen.PollenTypeCedar
	bestLevel := google_pollen.MinPollenLevel.ToInt()

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
