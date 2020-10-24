package db

import (
	"github.com/eakarpov/msaot/db/sqlite/migrations/msaot"
)

func InitDB() error {
	return msaot.MigrateSchema("test-db.db")
}
