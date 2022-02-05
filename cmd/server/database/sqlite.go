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

	if err = m.Down(); err != nil {
		fmt.Printf("migrate down error: %v \n", err)
		return err
	}


	fmt.Println("migrrate down done with success")

	return err
}

func (s *Sqlite) Exec(query string, args ...interface{}) (int64, error) {
	r, err := s.db.Exec(query, args...)
	if err != nil {
		return 0, err
	}
	if err != nil {
		return 0, fmt.Errorf("%s", err.Error())
	}

	return r.LastInsertId()
}

// Multi Select
func (s *Sqlite) Query(query string, dest interface{}, args ...interface{}) error {
	rows, err := s.db.Query(query, args...)
	if err != nil {
		return fmt.Errorf("%s", err.Error())
	}

	for rows.Next() {
		err := rows.Scan(&dest)
		if err != nil {
			return err
		}
	}

	return nil
}

// Single Select
func (s *Sqlite) QueryRow(query string, dest interface{}, args ...interface{}) error {
	row := s.db.QueryRow(query, args...)
	err := row.Scan(&dest)
	if err != nil {
    	if err == sql.ErrNoRows {
			return nil
    	} else {
			return err
    	}
	}
	return nil
}

func (s *Sqlite) Begin() (*Tx, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return nil, err
	}

	return &Tx{tx: tx}, nil
}

type Tx struct {
	tx          *sql.Tx
}

func (t *Tx) Exec(query string, args ...interface{}) (int64, error) {
	r, err := t.tx.Exec(query, args...)
	if err != nil {
		return 0, err
	}
	if err != nil {
		return 0, fmt.Errorf("%s", err.Error())
	}

	return r.LastInsertId()
}

// Multi Select
func (t *Tx) Query(query string, dest interface{}, args ...interface{}) error {
	rows, err := t.tx.Query(query, args...)
	if err != nil {
		return fmt.Errorf("%s", err.Error())
	}

	for rows.Next() {
		err := rows.Scan(&dest)
		if err != nil {
			return err
		}
	}

	return nil
}

// Single Select
func (t *Tx) QueryRow(query string, dest interface{}, args ...interface{}) error {
	row := t.tx.QueryRow(query, args...)
	err := row.Scan(&dest)
	if err != nil {
    	if err == sql.ErrNoRows {
			return nil
    	} else {
			return err
    	}
	}
	return nil
}

func (t *Tx) Commit() error {
	return t.tx.Commit()
}

func (t *Tx) Rollback() error {
	return t.tx.Rollback()
}
