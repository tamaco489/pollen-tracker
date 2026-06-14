package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/tamaco489/pollen-tracker/backend/internal/domain"
	"github.com/tamaco489/pollen-tracker/backend/internal/gen"
	"github.com/tamaco489/pollen-tracker/backend/pkg/errors/sentinel"
	"github.com/tamaco489/pollen-tracker/backend/pkg/library/google/pollen"
)

type getForecastUseCase struct {
	svc pollen.PollenService
}

func NewGetForecast(svc pollen.PollenService) GetForecastUseCase {
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
func (uc *getForecastUseCase) GetForecast(ctx context.Context, input gen.GetPollenParams) (*domain.PollenForecast, error) {
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

	// 常に最大日数 (maxDayOffset + 1 = 5) で取得して翌日以降の予報も得る
	request := &pollen.ForecastRequest{
		Lat:  input.Lat,
		Lng:  input.Lng,
		Days: maxDayOffset + 1,
	}

	// Google Pollen API から予報を取得
	resp, err := uc.svc.GetForecast(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("get forecast: %w", err)
	}

	// 対象日付の予報データが API レスポンスに含まれない場合はエラーを返す
	if len(resp.DailyForecasts) <= dayOffset {
		return nil, fmt.Errorf("google pollen api returned insufficient data: got %d days, expected index %d", len(resp.DailyForecasts), dayOffset)
	}

	pollenType, level := uc.dominantPollen(resp.DailyForecasts[dayOffset].Plants)

	// 指定日の翌日以降で API レスポンスに含まれる分を forecast として積む
	forecast := make([]domain.ForecastItem, 0, maxDayOffset-dayOffset)
	for i := dayOffset + 1; i < len(resp.DailyForecasts); i++ {
		_, forecastLevel := uc.dominantPollen(resp.DailyForecasts[i].Plants)
		forecast = append(forecast, domain.ForecastItem{
			Date:  today.Add(time.Duration(i) * 24 * time.Hour),
			Level: forecastLevel,
		})
	}

	return &domain.PollenForecast{
		Date:       targetDate,
		Level:      level,
		PollenType: pollenType,
		SeasonInfo: domain.SeasonCalendar[pollenType],
		Forecast:   forecast,
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
func (uc *getForecastUseCase) dominantPollen(plants []pollen.Plant) (domain.PollenType, int) {
	bestType := domain.PollenTypeCedar
	bestLevel := pollen.MinPollenLevel.ToInt()

	for _, p := range plants {
		pt, ok := domain.PlantCodeToType[p.Code]
		if !ok {
			continue
		}
		// 同じレベルなら InSeason の方を優先する (例: スギ花粉のピークは Level 3 だが、Level 2 でも InSeason ならピークに近いと判断する)
		if p.Level > bestLevel || (p.Level == bestLevel && p.InSeason) {
			bestType = pt
			bestLevel = p.Level
		}
	}

	return bestType, bestLevel
}
