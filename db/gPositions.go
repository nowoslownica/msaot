package db

import (
	"context"
	"database/sql"
	"github.com/eakarpov/msaot/db/sqlite/models"
)

type GPositions struct {
	db *sql.DB
}

func (c *sqliteDB) GPositions() *GPositions {
	return &GPositions{
		db: c.db,
	}
}

func (f *GPositions) GetAll(ctx context.Context) ([]*models.GrammarPosition, error) {
	return models.GrammarPositions().All(ctx, f.db)
}
