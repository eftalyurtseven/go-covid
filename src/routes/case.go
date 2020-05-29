package routes

import (
	"context"
	"fmt"
	"time"

	"github.com/eftalyurtseven/go-covid/src/config"
	"github.com/eftalyurtseven/go-covid/src/models"
	"github.com/eftalyurtseven/go-covid/src/utils"
	"github.com/gin-gonic/gin"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type GetParams struct {
	Date        string `json:"date"`
	CountryCode string `json:"country_code"`
}

func GetCases(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	var loginCmd GetParams
	c.BindJSON(&loginCmd)
	postDate := loginCmd.Date
	postCountry := loginCmd.CountryCode
	if len(postDate) == 0 || len(postCountry) == 0 {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "Please send date and country code!",
		})
		return
	}
	if len(postDate) != 10 {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "Invalid date format!",
		})
		return
	}

	if len(postCountry) < 2 {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "Invalid country code format!",
		})
		return
	}

	_, err := time.Parse("2006-01-02", postDate)
	if err != nil {
		fmt.Println("Err", err)
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "Date can't converted!",
		})
	}
	db, err := config.Connect()
	if err != nil {
		utils.SendSlackNotification(1, "DB Connection Err: "+err.Error())
		return
	}
	ctx := context.Background()
	check, err := models.Cases(
		qm.Where("countryterritoryCode = ?", postCountry),
	).One(ctx, db)
	if err != nil {
		fmt.Println(err)
	}
	if check == nil {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "Invalid country code!",
		})
		return
	}
	caseResult, err := models.Cases(
		qm.Where("countryterritoryCode = ? AND dateRep like ?", postCountry, "%"+postDate+"%"),
	).One(ctx, db)
	if err != nil {
		fmt.Println(err.Error())
	}
	if caseResult == nil {
		c.JSON(200, gin.H{
			"status":  "success",
			"cases":   nil,
			"message": "No affected row(s)!",
		})
		return
	}
	c.JSON(200, gin.H{
		"status": "success",
		"cases":  caseResult,
	})
	return
}
