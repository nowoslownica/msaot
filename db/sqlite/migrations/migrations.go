package migrations

import (
	"database/sql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	bindata "github.com/golang-migrate/migrate/v4/source/go_bindata"
)

//go:generate go run github.com/go-bindata/go-bindata/v3/go-bindata -pkg migrations -prefix "sqlite/migrations" -ignore .*\.go -o bindata.go ./...
//go:generate go fmt bindata.go

// MigrateSchema update db schema to last version saved with go-bindata
func MigrateSchema(conn string) error {
	sourceInstance, err := bindata.WithInstance(bindata.Resource(AssetNames(), Asset))
	if err != nil {
		return err
	}
	db, err := sql.Open("sqlite3", conn)
	if err != nil {
		return err
	}
	defer db.Close()
	targetInstance, err := sqlite3.WithInstance(db, new(sqlite3.Config))
	if err != nil {
		return err
	}
	defer targetInstance.Close()
	m, err := migrate.NewWithInstance("go-bindata", sourceInstance, "msaot", targetInstance)
	if err != nil {
		return err
	}
	// here error
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return err
	}
	return sourceInstance.Close()
}
