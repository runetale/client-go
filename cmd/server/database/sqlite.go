package database

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/Notch-Technologies/wizy/wislog"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
)

const SQLITE_DB_NAME = "wissy.db"

type SQLExecuter interface {
	Exec(query string, args ...interface{}) (int64, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
}

type Sqlite struct {
	db *sql.DB

	wislog *wislog.WisLog
}

func NewSqlite(wl *wislog.WisLog) *Sqlite {
	db, err := sql.Open("sqlite3", SQLITE_DB_NAME)
	if err != nil {
		panic(err)
	}

	return &Sqlite{
		db: db,

		wislog: wl,
	}
}

func (s *Sqlite) MigrationUp() error {
	driver, err := sqlite3.WithInstance(s.db, &sqlite3.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance("file://migrations", SQLITE_DB_NAME+"_foreign_keys=on?parseTime=true", driver)
	if err != nil {
		return err
	}

	if err = m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Printf("migrate up error: %v \n", err)
		return err
	}

	fmt.Println("migrate up done with success")

	return nil
}

func (s *Sqlite) MigrationDown() error {
	driver, err := sqlite3.WithInstance(s.db, &sqlite3.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance("file://migrations", SQLITE_DB_NAME+"_foreign_keys=off?parseTime=true", driver)
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

	return r.LastInsertId()
}

// Multi Select
func (s *Sqlite) Query(query string, args ...interface{}) (*sql.Rows, error) {
	rows, err := s.db.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("%s", err.Error())
	}

	return rows, nil

	//for rows.next() {
	//	err := rows.scan(&dest)
	//	if err != nil {
	//		return err
	//	}
	//}
}

// Single Select
func (s *Sqlite) QueryRow(query string, args ...interface{}) *sql.Row {
	row := s.db.QueryRow(query, args...)
	return row
}

func (s *Sqlite) Begin() (*Tx, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return nil, err
	}

	return &Tx{tx: tx}, nil
}

type Tx struct {
	tx *sql.Tx
}

func (t *Tx) Exec(query string, args ...interface{}) (int64, error) {
	r, err := t.tx.Exec(query, args...)
	if err != nil {
		return 0, err
	}
	return r.LastInsertId()
}

// Multi Select
func (t *Tx) Query(query string, args ...interface{}) (*sql.Rows, error) {
	rows, err := t.tx.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("%s", err.Error())
	}

	return rows, nil
}

// Single Select
func (t *Tx) QueryRow(query string, args ...interface{}) *sql.Row {
	row := t.tx.QueryRow(query, args...)
	return row
}

func (t *Tx) Commit() error {
	return t.tx.Commit()
}

func (t *Tx) Rollback() error {
	return t.tx.Rollback()
}
