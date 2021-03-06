package modules

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type UserComments struct {
	CommentId		int			`gorm:"primaryKey"`
	PostId			int			`json:"post_id"`
	Body			string		`json:"body"`
	CreatedOn		time.Time	`json:"created_on"`
	LastModified	time.Time	`json:"last_modified"`
	UserId 			int			`json:"user_id"`
}

type CreateUserComments struct {
	PostId			int			`json:"post_id" binding:"required"`
	Body			string		`json:"body"  binding:"required"`
	UserId 			int			`json:"user_id" binding:"required"`
}
type EditUserComments struct {
	CommentId		int			`json:"comment_id"`
	Body			string		`json:"body"`
	LastModified	time.Time	`json:"last_modified"`
}

func GetComments(c *gin.Context) {
	var comment []UserComments
	if err := db.Find(&comment).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, comment)
	}
}

func AddComments(c *gin.Context) {

	var input CreateUserComments
	//validate input
	if err := c.ShouldBindJSON(&input);err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//create Input
	comment := UserComments{UserId: input.UserId, PostId: input.PostId, Body: input.Body, CreatedOn: time.Now(), LastModified: time.Now()}
	if err := db.Create(&comment).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, comment)
	}
}

func GetCommentById(c *gin.Context) {
	id := c.Params.ByName("id")
	var comment UserComments
	if err := db.Where("comment_id = ?", id).First(&comment).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, comment)
	}
}

func GetCommentByPostId(c *gin.Context) {
	id := c.Params.ByName("id")
	var comment UserComments
	if err := db.Where("post_id = ?", id).Find(&comment).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, comment)
	}
}

func DelComment(c *gin.Context) {
	id := c.Params.ByName("id")
	var comment UserComments
	if err := db.Where("comment_id = ?", id).Delete(&comment).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, comment)
	}
}
func EditComment(c *gin.Context) {

	var input EditUserComments
	//validate input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//UPDATE DB
	comment := UserComments{Body: input.Body, CreatedOn: time.Now(), LastModified: time.Now()}
	if err := db.Model(&UserComments{}).Where("comment_id = ?", input.CommentId).Omit("created_on").Updates(&comment).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, "SUCCESS")
	}
}