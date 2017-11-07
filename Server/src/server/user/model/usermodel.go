package model

import (
	"database/sql"
	"errors"
	"server/user/model/user"
)

//RegisterUser returns a user-object if succsesfuly created in db
func RegisterUser(username string, password string, db *sql.DB) error {
	tx, err := db.Begin()
	stmt, err := db.Prepare("INSERT User SET username=?, password=?")
	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = stmt.Exec(username, password)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

//Login returns a user-object if found in db
func Login(username string, db *sql.DB) (*user.User, error) {
	stmt, err := db.Prepare("SELECT * FROM User WHERE username=?")
	if err != nil {
		return nil, err
	}
	res, err := stmt.Query(username)
	if err != nil {
		return nil, err
	}
	for res.Next() {
		var uid int
		var username string
		var password string
		err = res.Scan(&uid, &username, &password)
		if err != nil {
			return nil, err
		}
		return &user.User{uid, username, password}, nil
	}
	return nil, errors.New("No User")
}
