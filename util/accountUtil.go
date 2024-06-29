package util

import (
	"database/sql"

	"github.com/spending-tracking/db"
)

func CheckUsernameExist(username string) (bool, error) {

	err := db.QueryByUserName(username)

	if err != nil {
		if err != sql.ErrNoRows {
			return false, err
		}
		return false, nil
	}

	return true, nil
}

func CheckRawPassword(rawPassword, username string) (bool, error) {
	hash, err := db.GetHashPasswordByUserName(username)
	if err != nil {
		return false, err
	}

	return CheckPasswordHash(rawPassword, hash), nil
}
