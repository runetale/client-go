package database

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
)

const SQLITE_DB_NAME = "wissy.db"

type Sqlite struct {
	db *sql.DB
}

func NewSqlite() *Sqlite {
	db, err := sql.Open("sqlite3", SQLITE_DB_NAME)
	if err != nil {
		panic(err)
	}

	return &Sqlite{
		db: db,
	}
}

func (s *Sqlite) MigrationUp() error {
	driver, err := sqlite3.WithInstance(s.db, &sqlite3.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance("file://migrations", SQLITE_DB_NAME + "_foreign_keys=on", driver)
	if err != nil {
		return err
	}

	if err = m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Printf("migrate up error: %v \n", err)
		return err
	}

	fmt.Println("migrrate up done with success")

	return err
}

func (s *Sqlite) MigrationDown() error {
	driver, err := sqlite3.WithInstance(s.db, &sqlite3.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance("file://migrations", SQLITE_DB_NAME + "_foreign_keys=off", driver)
	if err != nil {
		return err
	}

	if err = m.Down(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Printf("migrate down error: %v \n", err)
		return err
	}

	fmt.Println("migrrate down done with success")

	return err
}
