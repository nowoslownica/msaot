package db

import (
	"database/sql"
	migrations "github.com/eakarpov/msaot/db/sqlite/migrations"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type CloseFunc func() error

type sqliteDB struct {
	db  *sql.DB
	name string
}

func New() *sqliteDB {
	db := &sqliteDB{
		db:  nil,
		name: "test-db",
	}
	return db
}

func (c *sqliteDB) Init() error {
	db, err := sql.Open("sqlite3", c.name)
	if err != nil {
		return err
	}
	boil.SetDB(db)
	err = db.Ping()
	if err != nil {
		return err
	}
	err = InitDB()
	if err != nil {
		return err
	}
	c.db = db
	return nil
}

func (c *sqliteDB) Close() error {
	return c.db.Close()
}

func InitDB() error {
	return migrations.MigrateSchema("test-db.db")
}
