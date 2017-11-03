package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	db *sql.DB
}

func (database *Database) Connect(connection string, username string, password string) error {
	db, err := sql.Open("mysql", username+":"+password+"@/"+connection)
	if err != nil {
		return err
	}
	err = db.Ping()
    if err != nil {
		return err
	}
	database.db = db
	return nil
}

func (database Database) Close() {
	database.Close()
}