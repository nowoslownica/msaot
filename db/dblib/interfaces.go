package dblib

import (
	"context"
	"github.com/eakarpov/msaot/db/sqlite/models"
)

type DB interface {
	Flexies() *Flexies
}

type Flexies interface {
	GetByValue(ctx context.Context, word string) (values []*models.Flexy, err error)
}
