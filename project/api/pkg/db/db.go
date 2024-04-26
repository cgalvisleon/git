package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

func ConectDB() error {
	log.Println("Connecting to database")

	if Db != nil {
		return nil
	}

	var err error
	Db, err = sql.Open("sqlte3", "mydb.db")
	if err != nil {
		return err
	}

	err = Db.Ping()
	if err != nil {
		return err
	}

	log.Panicln("Connected to database successfully")

	return nil
}
