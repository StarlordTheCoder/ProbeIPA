package model

import (
	"database/sql"
	"server/user/model/user"
)

func RegisterUser(username string, password string, db *sql.DB) user.User {
	stmt, err := db.Prepare("INSERT User SET username=?, password?")
	checkError(err)
	_, err = stmt.Exec(username, password)
	checkError(err)
	
	return Login(username, db)
}

func Login(username string, db *sql.DB) user.User {
	stmt, err := db.Prepare("SELECT * FROM User WHERE username=?")
	checkError(err)
	res, err := stmt.Query(username)
	checkError(err)
	for res.Next() {
		var uid int
		var username string
		var password string
		err = res.Scan(&uid, &username, &password)
		return user.User{uid, username, password}
	}
	return user.User{nil, nil, nil}
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
