package db

import (
	"context"
	"database/sql"
	"github.com/eakarpov/msaot/db/sqlite/models"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type Flexies struct {
	db *sql.DB
}

func (c *sqliteDB) Flexies() *Flexies {
	return &Flexies{
		db: c.db,
	}
}

func (f *Flexies) GetByValue(ctx context.Context, word string) ([]*models.Flexy, error) {
	return models.Flexies(models.FlexyWhere.Value.EQ(word), qm.Load(models.FlexyRels.GPositionGrammarPosition)).All(ctx, f.db)
}
