package datastore

import (
	"github.com/tamaco489/pollen-tracker/backend/internal/infrastructure/datastore/gen"
	"github.com/tamaco489/pollen-tracker/backend/pkg/infrastructure/datastore"
)

type symptomsRepository struct {
	db      *datastore.DB
	queries *gen.Queries
}

func NewSymptomsRepository(db *datastore.DB) *symptomsRepository {
	return &symptomsRepository{
		db:      db,
		queries: gen.New(),
	}
}
