package datastore

import "context"

func (r *symptomsRepository) GetSymptomPollenLevels(ctx context.Context) ([]int64, error) {
	return r.queries.GetSymptomPollenLevels(ctx, r.db)
}
