package main

//go:generate sqlboiler --wipe mysql

import (
	"fmt"

	"github.com/eftalyurtseven/go-covid/src/routes"
	"github.com/eftalyurtseven/go-covid/src/utils"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	fmt.Println("Hello, please select a type:")
	fmt.Println("1 - Excel to mysql")
	fmt.Println("2 - API Server")
	var pType int
	fmt.Scanf("%d", &pType)

	if pType == 1 {
		utils.Insert()
	} else {
		router := gin.Default()
		router.POST("/cases", routes.GetCases)
		router.Run()
	}

	//utils.Insert()

}
