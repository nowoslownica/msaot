package dictionary

import (
	"context"
	"github.com/eakarpov/msaot/db"
	"github.com/eakarpov/msaot/db/dblib"
	"github.com/eakarpov/msaot/db/sqlite/models"
)

func InitMain() (dblib.DB, db.CloseFunc, error) {
	mainDB := db.New()
	err := mainDB.Init()
	if err != nil {
		return nil, nil, err
	}
	dblib.Database = mainDB
	return mainDB, mainDB.Close, nil
}

func GetFlexForm(word string) ([]*models.Flexy, error) {
	flexForm, err := dblib.Database.Flexies().GetByValue(context.Background(), word)
	if err != nil {
		return nil, err
	}
	return flexForm, nil
}
