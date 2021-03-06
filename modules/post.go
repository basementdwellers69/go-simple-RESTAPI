package modules

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

type UserPosts struct {
	PostId			int			`gorm:"primaryKey"`
	UserId 			int			`json:"user_id"`
	Body			string		`json:"body"`
	Likes			int			`json:"likes"`
	CreatedOn		time.Time	`json:"created_on"`
	LastModified	time.Time	`json:"last_modified"`
}

type CreateUserPosts struct {
	UserId 			int			`json:"user_id" binding:"required"`
	Body			string		`json:"body"  binding:"required"`
}
type EditUserPosts struct {
	PostId			string		`json:"post_id"`
	Body			string		`json:"body"`
	LastModified	time.Time	`json:"last_modified"`
}

func GetPost(c *gin.Context) {
	var post []UserPosts
	if err := db.Find(&post).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, post)
	}
}

func AddPost(c *gin.Context) {

	var input CreateUserPosts
	//validate input
	if err := c.ShouldBindJSON(&input);err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//create Input
	post:= UserPosts{UserId:input.UserId, Body: input.Body, Likes: 0, CreatedOn: time.Now(), LastModified: time.Now()}
	if err := db.Create(&post).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, post)
	}
}

func GetPostById(c *gin.Context) {
	id := c.Params.ByName("id")
	var comment UserComments
	if err := db.Where("post_id = ?", id).First(&comment).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, comment)
	}
}

func DelPost(c *gin.Context) {
	id := c.Params.ByName("id")
	var post UserPosts
	if err := db.Where("comment_id = ?", id).Delete(&post).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, post)
	}
}
func EditPost(c *gin.Context) {
	//GET OMIT TYPES
	temp := c.Params.ByName("type")
	types := strings.Split(temp, ",")

	var input EditUserPosts
	//validate input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//UPDATE DB
	post := UserPosts{Body: input.Body, Likes:0, CreatedOn: time.Now(), LastModified: time.Now()}
	if len(types) == 1 {
		if err := db.Model(&UserPosts{}).Where("post_id = ?", input.PostId).Omit(types[0],"created_on").Updates(&post).Error; err != nil {
			c.AbortWithStatus(404)
			fmt.Println(err)
		} else {
			c.JSON(200, "SUCCESS")
		}
	}else{
		c.JSON(200, "ERROR")
	}
}