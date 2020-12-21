package dblib

import (
	"context"
	"github.com/eakarpov/msaot/db"
	"github.com/eakarpov/msaot/db/sqlite/models"
	"github.com/eakarpov/msaot/lexicon/lemmas"
	"github.com/volatiletech/null/v8"
)

var Database DB

type DB interface {
	Flexies() *db.Flexies
	Lemmas() *db.Lemmas
	GPositions() *db.GPositions
}

type GPositions interface {
	GetAll(ctx context.Context) (values []*models.GrammarPosition, err error)
}

type Flexies interface {
	GetByValue(ctx context.Context, word string) (values []*models.Flexy, err error)
	AddOne(ctx context.Context, lemma *lemmas.Lemma, lemmaId null.Int64, gPosId null.Int64) (models.Flexy, error)
}

type Lemmas interface {
	AddOne(ctx context.Context, lemma *lemmas.Lemma) (models.Lemma, error)
}
