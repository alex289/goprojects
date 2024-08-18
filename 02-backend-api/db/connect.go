package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func getConnection() (*sql.DB, error) {
	connStr := "user=postgres password=postgres dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
