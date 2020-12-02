package lexicon

import (
	"github.com/eakarpov/msaot/db"
	"github.com/eakarpov/msaot/db/dblib"
)

func initMain() (dblib.DB, db.CloseFunc, error) {
	mainDB := db.New()
	err := mainDB.Init()
	if err != nil {
		return nil, nil, err
	}
	return mainDB, mainDB.Close, nil
}

func GetFlexForm(word string)  {
}
