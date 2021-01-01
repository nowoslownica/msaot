package db

import (
	"database/sql"
	"github.com/eakarpov/msaot/db/sqlite/migrations/full"
	"github.com/eakarpov/msaot/db/sqlite/migrations/structure"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"path/filepath"
)

type CloseFunc func() error

type sqliteDB struct {
	db  *sql.DB
	name string
}

func New() *sqliteDB {
	db := &sqliteDB{
		db:  nil,
		name: filepath.Join("../test-db.db"),
	}
	return db
}

func GetByName(name string) *sqliteDB {
	db := &sqliteDB{
		db:  nil,
		name: filepath.Join(name),
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
	err = c.InitDB(true)
	if err != nil {
		return err
	}
	c.db = db
	return nil
}

func (c *sqliteDB) Close() error {
	return c.db.Close()
}

func (c *sqliteDB) InitDB(data bool) error {
	if data {
		return full.MigrateSchema(c.name)
	}
	return structure.MigrateSchema(c.name)
}
