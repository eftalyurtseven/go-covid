package routes

import (
	"context"
	"regexp"

	"github.com/eftalyurtseven/go-covid/src/config"
	"github.com/eftalyurtseven/go-covid/src/models"
	"github.com/eftalyurtseven/go-covid/src/utils"
	"github.com/gin-gonic/gin"
	guuid "github.com/google/uuid"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type BodyParams struct {
	Email string `json:"email"`
}

func generateUID() string {
	id := guuid.New()
	return id.String()
}

// Register func
func Register(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	var parameters = new(BodyParams)
	err := c.BindJSON(&parameters)
	if err != nil {
		utils.SendSlackNotification(1, "/auth/register JSON Bind error"+err.Error())
		c.JSON(400, gin.H{
			"status": "json_parse_error",
		})
		return
	}
	email := parameters.Email
	if len(email) == 0 {
		c.JSON(400, gin.H{
			"status": "missing_parameter",
		})
		return
	}
	mailCheck := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if mailCheck.MatchString(email) == false {
		c.JSON(400, gin.H{
			"status": "invalid_email",
		})
		return
	}
	db, err := config.Connect()
	if err != nil {
		utils.SendSlackNotification(1, "DB Connection Err: "+err.Error())
		c.JSON(500, gin.H{
			"status": "server_error",
		})
		return
	}
	defer db.Close()
	check, err := models.Users(
		qm.Where("email = ?", email),
	).One(context.Background(), db)
	if err != nil {
		utils.SendSlackNotification(1, "/auth/register check mail error")
		c.JSON(500, gin.H{
			"status": "server_error",
		})
		return
	}

	if check != nil {
		c.JSON(209, gin.H{
			"status": "registered_user",
		})
		return
	}
	apiKey := generateUID()
	var userModel = new(models.User)
	userModel.Email = email
	userModel.APIKey = apiKey
	userModel.IP = c.ClientIP()
	err = userModel.Insert(context.Background(), db, boil.Infer())
	if err != nil {
		utils.SendSlackNotification(1, "/auth/register insert error!")
		c.JSON(500, gin.H{
			"status": "server_error",
		})
		return
	}

	c.JSON(200, gin.H{
		"status":  "success",
		"api_key": apiKey,
	})

	return
}
