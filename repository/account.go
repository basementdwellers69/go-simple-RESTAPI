package repository

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sqlcTest/config"
	"sqlcTest/domain"
	"sqlcTest/dto"
	"strings"
	"time"
)

var db = config.GetConnection()

func GetAccount(c *gin.Context) {
	var accounts []domain.Accounts
	if err := db.Find(&accounts).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, accounts)
	}
}

func AddAccount(c *gin.Context) {

	var input dto.CreateAccounts
	//validate input
	if err := c.ShouldBindJSON(&input);err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//create Input
	account := domain.Accounts{Username: input.Username, Password: input.Password, Email: input.Email,CreatedOn: time.Now(), LastLogin: time.Now()}
	if err := db.Create(&account).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, account)
	}
}

func GetAccountId(c *gin.Context) {
	id := c.Params.ByName("id")
	var account domain.Accounts
	if err := db.Where("user_id = ?", id).First(&account).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, account)
	}
}

func DelAccount(c *gin.Context) {
	id := c.Params.ByName("id")
	var account domain.Accounts
	if err := db.Where("user_id = ?", id).Delete(&account).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, account)
	}
}
func EditAccount(c *gin.Context) {
	//GET OMIT TYPES
	temp := c.Params.ByName("type")
	types :=	strings.Split(temp, ",")

	var input dto.EditAccounts
	//validate input
	if err := c.ShouldBindJSON(&input);err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//UPDATE DB
	account := domain.Accounts{Username: input.Username, Password: input.Password, Email: input.Email,CreatedOn: time.Now(), LastLogin: time.Now()}
	if len(types) == 1{
		if err := db.Model(&domain.Accounts{}).Where("user_id = ?", input.ID).Omit(types[0], "created_on").Updates(&account).Error; err != nil {
			c.AbortWithStatus(404)
			fmt.Println(err)
		} else {
			c.JSON(200, "SUCCESS")
		}
		return
	}else if len(types) == 2{
		if err := db.Model(&domain.Accounts{}).Where("user_id = ?", input.ID).Omit(types[0],types[1], "created_on").Updates(&account).Error; err != nil {
			c.AbortWithStatus(404)
			fmt.Println(err)
		} else {
			c.JSON(200, "SUCCESS")
		}
		return
	}else{
		c.JSON(200, "ERROR OCCURRED")
	}
}