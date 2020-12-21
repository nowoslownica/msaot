package db

import (
	"context"
	"database/sql"
	"github.com/eakarpov/msaot/db/sqlite/models"
	"github.com/eakarpov/msaot/lexicon/lemmas"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
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
	return models.Flexies(
		models.FlexyWhere.Value.EQ(word),
		qm.Load(models.FlexyRels.GPositionIdGrammarPosition),
		qm.Load(models.FlexyRels.LemmaIdLemmas),
	).All(ctx, f.db)
}

func (f *Flexies) AddOne(
	ctx context.Context,
	flexy *lemmas.Lemma,
	lemmaId null.Int64,
	gPosId null.Int64,
) (models.Flexy, error) {
	var fl1 models.Flexy
	fl1.Value = flexy.Value
	fl1.LemmaId = lemmaId
	fl1.GPositionId = gPosId
	err := fl1.Insert(ctx, f.db, boil.Infer())
	return fl1, err
}
