package main

//go:generate sqlboiler --wipe mysql

import (
	"fmt"
	"log"

	"github.com/eftalyurtseven/go-covid/src/routes"
	"github.com/eftalyurtseven/go-covid/src/utils"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/mileusna/crontab"
)

func main() {

	ctab := crontab.New() // create cron table

	// AddJob and test the errors
	err := ctab.AddJob("* * * * *", cronFunc) // on 1st day of month
	if err != nil {
		log.Println(err)
		return
	}
	router := gin.Default()
	router.POST("/cases", routes.GetCases)
	router.Run()

	/*
		if pType == 1 {
			utils.Insert()
		} else {
			router := gin.Default()
			router.POST("/cases", routes.GetCases)
			router.Run()
		}
	*/

	//utils.Insert()

}
func cronFunc() {
	fmt.Println("Cron started!")
	err := utils.SendSlackNotification(2, "Cron started!")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = utils.DownloadFile("covid.xlsx", "https://www.ecdc.europa.eu/sites/default/files/documents/COVID-19-geographic-disbtribution-worldwide.xlsx")
	if err != nil {
		fmt.Println(err.Error())
	}
	utils.Insert()
	defer utils.SendSlackNotification(2, "Cron ended!")

}
