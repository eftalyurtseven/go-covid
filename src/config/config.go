package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {
	databaseHost := "DATABASE_HOST"
	databasePort := "DATABASE_PORT"
	databaseName := "DB_NAME"
	databaseUser := "DB_USER"
	databasePass := "DB_PASS"
	db, err := sql.Open("mysql", databaseUser+":"+databasePass+"@tcp("+databaseHost+":"+databasePort+")/"+databaseName+"?parseTime=true")
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
