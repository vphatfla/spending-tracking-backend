package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func OpenDB() *sql.DB {
	var err error

	db, err = sql.Open("mysql", "test_user:password@tcp(localhost:3307)/test_db")

	if err != nil {
		fmt.Printf("ERROR opening db connection %e\n", err)
		return nil
	}

	fmt.Printf("Successfully connect to mysql db\n")
	return db
}

func GetDBConn() *sql.DB {
	if db == nil {
		return OpenDB()
	}

	return db
}
