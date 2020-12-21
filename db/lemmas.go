package db

import (
	"context"
	"database/sql"
	"github.com/eakarpov/msaot/db/sqlite/models"
	"github.com/eakarpov/msaot/lexicon/gender"
	"github.com/eakarpov/msaot/lexicon/lemmas"
	"github.com/eakarpov/msaot/lexicon/number"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type Lemmas struct {
	db *sql.DB
}

func (c *sqliteDB) Lemmas() *Lemmas {
	return &Lemmas{
		db: c.db,
	}
}

func (l *Lemmas) getGrammarPosition(ctx context.Context, lemma *lemmas.Lemma) (*models.GrammarPosition, error) {
	genderValue := null.Int64{
		Int64: 0,
		Valid: false,
	}
	numberValue := null.Int64{
		Int64: 0,
		Valid: false,
	}
	if lemma.NCase != nil {
		if lemma.NCase.Number != number.Number(0) {
			numberValue.Int64 = int64(lemma.NCase.Number)
			numberValue.Valid = true
		}
		if lemma.NCase.Gender != gender.Gender(0) {
			genderValue.Int64 = int64(lemma.NCase.Gender)
			genderValue.Valid = true
		}
	}
	if lemma.VCase != nil {
		if lemma.VCase.Number != number.Number(0) {
			numberValue.Int64 = int64(lemma.VCase.Number)
			numberValue.Valid = true
		}
		if lemma.VCase.Gender != gender.Gender(0) {
			genderValue.Int64 = int64(lemma.VCase.Gender)
			genderValue.Valid = true
		}
	}
	return models.GrammarPositions(
		models.GrammarPositionWhere.GGender.EQ(genderValue),
		models.GrammarPositionWhere.GNumber.EQ(numberValue),
	).One(ctx, l.db)
}

func (l *Lemmas) AddOne(ctx context.Context, lemma *lemmas.Lemma) (models.Lemma, error) {
	var l1 models.Lemma
	l1.Pos = string(lemma.Pos)
	l1.Value = lemma.Normal
	if lemma.Indeclinable {
		l1.ChangeSchema = null.Int64{
			Int64: 0,
			Valid: true,
		}
	}

	err := l1.Insert(ctx, l.db, boil.Infer())
	return l1, err
}