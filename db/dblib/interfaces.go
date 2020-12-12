package dblib

import (
	"context"
	"github.com/eakarpov/msaot/db"
	"github.com/eakarpov/msaot/db/sqlite/models"
	"github.com/eakarpov/msaot/lexicon/lemmas"
)

var Database DB

type DB interface {
	Flexies() *db.Flexies
	Lemmas() *db.Lemmas
}

type Flexies interface {
	GetByValue(ctx context.Context, word string) (values []*models.Flexy, err error)
}

type Lemmas interface {
	AddLemma(ctx context.Context, lemma *lemmas.Lemma) (error)
}
