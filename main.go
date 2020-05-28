package main

//go:generate sqlboiler --wipe mysql

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/eftalyurtseven/go-covid/src/config"
	"github.com/eftalyurtseven/go-covid/src/models"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func StrToInt(str string) (int, error) {
	if len(str) == 0 {
		return 0, nil
	}
	nonFractionalPart := strings.Split(str, ".")
	return strconv.Atoi(nonFractionalPart[0])
}

func insert() {
	log.Println("Program started")
	defer log.Println("Ended")
	db, err := config.Connect()
	if err != nil {
		panic(err)
	}
	f, err := excelize.OpenFile("covid.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}

	rows := f.GetRows("COVID-19-geographic-disbtributi")
	for index, row := range rows {
		if index == 0 {
			continue
		}
		day, err := StrToInt(row[1])
		if err != nil {
			panic(err)
		}

		month, err := StrToInt(row[2])
		if err != nil {
			panic(err)
		}

		year, err := StrToInt(row[3])
		if err != nil {
			panic(err)
		}

		cases, err := StrToInt(row[4])
		if err != nil {
			panic(err)
		}

		deaths, err := StrToInt(row[5])
		if err != nil {
			panic(err)
		}

		popData, err := StrToInt(row[9])
		if err != nil {
			panic(err)
		}
		dayFormatted := row[1]
		monthFormatted := row[2]
		if len(dayFormatted) == 1 {
			dayFormatted = "0" + dayFormatted
		}
		if len(monthFormatted) == 1 {
			monthFormatted = "0" + monthFormatted
		}

		dateRep := row[3] + "-" + monthFormatted + "-" + dayFormatted

		var caseModel models.Case
		caseModel.DateRep = dateRep
		caseModel.Day = day
		caseModel.Month = month
		caseModel.Year = year
		caseModel.Cases = cases
		caseModel.Deaths = deaths
		caseModel.CountriesAndTerritories = row[6]
		caseModel.GeoID = row[7]
		caseModel.CountryterritoryCode = row[8]
		caseModel.PopData2018 = popData
		caseModel.ContinentExp = row[10]

		caseModel.Insert(context.Background(), db, boil.Infer())

	}
}

func main() {
	db, err := config.Connect()
	if err != nil {
		panic(err)
	}
	r := gin.Default()
	r.GET("/cases", func(c *gin.Context) {
		date := c.Query("date")
		country := c.Query("country")
		fmt.Println(date)
		fmt.Println(country)

		ctx := context.Background()
		cases, err := models.Cases(
			qm.Where("countryterritoryCode = ? AND dateRep like ?", country, "%"+date+"%"),
		).One(ctx, db)
		if err != nil {
			fmt.Println(err)
		}

		c.JSON(200, gin.H{
			"message": "success",
			"result":  cases,
		})
	})
	r.Run()
}
