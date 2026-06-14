package usecase

import (
	"context"
	"fmt"
	"slices"

	"github.com/tamaco489/pollen-tracker/backend/internal/gen"
	"github.com/tamaco489/pollen-tracker/backend/internal/infrastructure/datastore"
)

type getThresholdUseCase struct {
	repo datastore.GetThresholdRepository
}

func NewGetThreshold(repo datastore.GetThresholdRepository) GetThresholdUseCase {
	return &getThresholdUseCase{repo: repo}
}

func (uc *getThresholdUseCase) GetThreshold(ctx context.Context) (*gen.ThresholdResponse, error) {
	levels, err := uc.repo.GetSymptomPollenLevels(ctx)
	if err != nil {
		return nil, fmt.Errorf("get symptom pollen levels: %w", err)
	}

	// fallbackThreshold はサンプルデータが0件のときに返すデフォルト値 (5段階中の中間値)
	const fallbackThreshold int32 = 3
	if len(levels) == 0 {
		return &gen.ThresholdResponse{
			Threshold:   fallbackThreshold,
			SampleCount: 0,
			IsEstimated: true,
		}, nil
	}

	slices.Sort(levels)

	return &gen.ThresholdResponse{
		Threshold:   uc.medianOf(levels),
		SampleCount: int32(len(levels)),
		IsEstimated: false,
	}, nil
}

// medianOf はソート済みの花粉レベル一覧から中央値を返す
//
// 奇数件数: 中央の値をそのまま返す
// 偶数件数: 中央 2 値の平均を整数で返す (切り捨て)
func (uc *getThresholdUseCase) medianOf(sorted []int64) int32 {
	n := len(sorted)
	if n%2 == 1 {
		return uc.pickMiddle(sorted, n)
	}
	return uc.averageOfMiddlePair(sorted, n)
}

// pickMiddle は奇数長スライスの中央要素を返す
func (uc *getThresholdUseCase) pickMiddle(sorted []int64, n int) int32 {
	return int32(sorted[n/2])
}

// averageOfMiddlePair は偶数長スライスの中央 2 値の平均を返す
func (uc *getThresholdUseCase) averageOfMiddlePair(sorted []int64, n int) int32 {
	return int32((sorted[n/2-1] + sorted[n/2]) / 2)
}
