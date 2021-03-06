package modules

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sqlcTest/conn"
	"strings"
	"time"
)

type Accounts struct {
	UserId		int			`gorm:"primaryKey"`
	Username	string		`json:"username"`
	Password	string		`json:"password"`
	Email		string		`json:"email"`
	CreatedOn	time.Time	`json:"created_on"`
	LastLogin	time.Time	`json:"last_login"`
}

type CreateAccounts struct {
	Username	string		`json:"username" binding:"required"`
	Password	string		`json:"password" binding:"required"`
	Email		string		`json:"email" binding:"required"`
}
type EditAccounts struct {
	ID			int			`json:"user_id"`
	Username	string		`json:"username"`
	Password	string		`json:"password"`
	Email		string		`json:"email"`
}

var db = conn.GetConnection()

func GetAccount(c *gin.Context) {
	var accounts []Accounts
	if err := db.Find(&accounts).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, accounts)
	}
}

func AddAccount(c *gin.Context) {

	var input CreateAccounts
	//validate input
	if err := c.ShouldBindJSON(&input);err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//create Input
	account := Accounts{Username: input.Username, Password: input.Password, Email: input.Email,CreatedOn: time.Now(), LastLogin: time.Now()}
	if err := db.Create(&account).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, account)
	}
}

func GetAccountId(c *gin.Context) {
	id := c.Params.ByName("id")
	var account Accounts
	if err := db.Where("user_id = ?", id).First(&account).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, account)
	}
}

func DelAccount(c *gin.Context) {
	id := c.Params.ByName("id")
	var account Accounts
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

	var input EditAccounts
	//validate input
	if err := c.ShouldBindJSON(&input);err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//UPDATE DB
	account := Accounts{Username: input.Username, Password: input.Password, Email: input.Email,CreatedOn: time.Now(), LastLogin: time.Now()}
	if len(types) == 1{
		if err := db.Model(&Accounts{}).Where("user_id = ?", input.ID).Omit(types[0], "created_on").Updates(&account).Error; err != nil {
			c.AbortWithStatus(404)
			fmt.Println(err)
		} else {
			c.JSON(200, "SUCCESS")
		}
		return
	}else if len(types) == 2{
		if err := db.Model(&Accounts{}).Where("user_id = ?", input.ID).Omit(types[0],types[1], "created_on").Updates(&account).Error; err != nil {
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