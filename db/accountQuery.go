package db

import (
	"database/sql"
	"fmt"

	"github.com/spending-tracking/model"
)

func GetAccountById(id int) (*model.User, error) {
	db := GetDBConn()

	query := "SELECT id, username, name FROM user WHERE id = ?"
	row := db.QueryRow(query, id)
	var user model.User
	if err := row.Scan(&user.ID, &user.Username, &user.Name); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user with id = %d not found", id)
		}
		return nil, err
	}

	return &user, nil
}
