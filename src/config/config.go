package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {
	databaseHost := "qzkp8ry756433yd4.cbetxkdyhwsb.us-east-1.rds.amazonaws.com"
	databasePort := "3306"
	databaseName := "n5mt92h8iv9gvg7i"
	databaseUser := "ct49wr3k1p0a8gis"
	databasePass := "kmqgdy4auqs1ln5p"
	db, err := sql.Open("mysql", databaseUser+":"+databasePass+"@tcp("+databaseHost+":"+databasePort+")/"+databaseName+"?parseTime=true")
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
