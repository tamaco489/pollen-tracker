package datastore

import (
	"github.com/tamaco489/pollen-tracker/backend/internal/infrastructure/datastore/gen"

	pkg_datastore "github.com/tamaco489/pollen-tracker/backend/pkg/infrastructure/datastore"
)

type symptomsRepository struct {
	db      *pkg_datastore.DB
	queries *gen.Queries
}

func NewSymptomsRepository(db *pkg_datastore.DB) *symptomsRepository {
	return &symptomsRepository{db: db, queries: gen.New()}
}
