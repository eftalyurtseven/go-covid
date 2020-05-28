package main

//go:generate sqlboiler --wipe mysql

import (
	"log"

	"github.com/eftalyurtseven/go-covid/src/config"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	_, err := config.Connect("root:@tcp(127.0.0.1:3306)/demo?parseTime=true")
	if err != nil {
		log.Fatalf("failed to mysql open %+v", err)
	}

	// set global database
	// boil.SetDB(db)

	/*
		u := &models.Case{
			DateRep:                 "test",
			Day:                     1,
			Month:                   1,
			Year:                    1,
			Cases:                   1,
			Deaths:                  1,
			CountriesAndTerritories: "1",
			GeoID:                   "1",
			CountryterritoryCode:    "1",
			PopData2018:             1,
			ContinentExp:            "1",
		}
		var tx *sql.Tx
		var ctx context.Context
		ctx, _ = context.WithTimeout(context.Background(), 15*time.Second)
		tx, err = db.BeginTx(ctx, nil)
		err1 := u.Insert(context.Background(), tx, boil.Infer())
		if err1 != nil {
			fmt.Println(err1)
		}
		fmt.Println("db")
		fmt.Println("inserted")
	*/
}
