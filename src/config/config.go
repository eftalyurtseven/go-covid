package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {

	databaseHost := "127.0.0.1"
	databasePort := "3306"
	databaseName := "covid"
	databaseUser := "root"
	databasePass := ""
	db, err := sql.Open("mysql", databaseUser+":"+databasePass+"@tcp("+databaseHost+":"+databasePort+")/"+databaseName+"?parseTime=true")
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
