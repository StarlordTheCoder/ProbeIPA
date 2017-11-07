package authorisation

import (
	"database/sql"
	"server/bcrypt"
	"server/user/model"
)

//Authorisate checks authorisation
func Authorisate(username string, password string, db *sql.DB) (*int, error) {
	user, err := model.Login(username, db)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashes([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, err
	}
	return &user.Id, nil
}
