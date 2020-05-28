package config

import (
	// Import this so we don't have to use qm.Limit etc.
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/volatiletech/sqlboiler/boil"
)

func GetDB() (*sql.DB, error) {
	// Open handle to database like normal
	dbhost := "localhost"
	dbport := "3306"
	dbname := "covid"
	dbusername := "root2"
	dbpassword := ""

	db, err := sql.Open("mysql", dbusername+":"+dbpassword+"@tcp(["+dbhost+"]:"+dbport+")/"+dbname)
	if err != nil {
		fmt.Println(err)
	}

	// If you don't want to pass in db to all generated methods
	// you can use boil.SetDB to set it globally, and then use
	// the G variant methods like so (--add-global-variants to enable)
	boil.SetDB(db)
	boil.DebugMode = true
	return db, err
}
