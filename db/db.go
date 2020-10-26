package db

import (
	"github.com/eakarpov/msaot/db/sqlite/migrations"
)

func InitDB() error {
	return migrations.MigrateSchema("test-db.db")
}
