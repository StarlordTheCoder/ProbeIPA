package controler

import (
	"database/sql"
	"server/bcrypt"
	"server/user/model"
)

//RegisterUser registers a user
func RegisterUser(username string, password string, db *sql.DB) error {
	hashedPassword, err := bcrypt.Encrypt([]byte(password))
	if err != nil {
		return err
	}
	err = model.RegisterUser(username, string(hashedPassword), db)
	if err != nil {
		return err
	}
	return nil
}
