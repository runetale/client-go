package database

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

const SQLITE_DB_NAME = "wissy.db"

type Sqlite struct {

}

func NewSqlite() {
	db, err := sql.Open("sqlite3", SQLITE_DB_NAME)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(db)
}
