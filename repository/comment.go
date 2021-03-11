package repository

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sqlcTest/domain"
	"sqlcTest/dto"
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

type Result struct{
	CommentId		int			`gorm:"primaryKey"`
	PostId			int			`json:"post_id"`
	Username 		string		`json:"username"`
	Body			string		`json:"body"`
	CreatedOn		time.Time	`json:"created_on"`
	LastModified	time.Time	`json:"last_modified"`
}

func GetComments(c *gin.Context) {
	var comment []domain.UserComments
	var res []dto.CommentsResponse
	if err := db.Model(&comment).Select("comment_id ,post_id, accounts.username, body, user_comments.created_on, user_comments.last_modified").Joins("INNER JOIN accounts ON user_comments.user_id = accounts.user_id").Find(&res).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, res)
	}
}

func AddComments(c *gin.Context) {

	var input dto.CreateUserComments
	//validate input
	if err := c.ShouldBindJSON(&input);err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//create Input
	comment := domain.UserComments{UserId: input.UserId, PostId: input.PostId, Body: input.Body, CreatedOn: time.Now(), LastModified: time.Now()}
	if err := db.Create(&comment).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, comment)
	}
}

func GetCommentById(c *gin.Context) {
	id := c.Params.ByName("id")
	var comment domain.UserComments
	var res []dto.CommentsResponse
	if err := db.Model(&comment).Select("comment_id ,post_id, body, user_comments.created_on, user_comments.last_modified, accounts.username").Joins("INNER JOIN accounts ON user_comments.user_id = accounts.user_id").Where("comment_id = ?", id).First(&res).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, res)
	}
}

func GetCommentByPostId(c *gin.Context) {
	id := c.Params.ByName("id")
	var comment domain.UserComments
	var res []dto.CommentsResponse
	if err := db.Model(&comment).Select("comment_id ,post_id, body, user_comments.created_on, user_comments.last_modified, accounts.username").Joins("INNER JOIN accounts ON user_comments.user_id = accounts.user_id").Where("post_id = ?", id).Find(&res).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, res)
	}
}

func DelComment(c *gin.Context) {
	id := c.Params.ByName("id")
	var comment domain.UserComments
	if err := db.Where("comment_id = ?", id).Delete(&comment).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, comment)
	}
}
func EditComment(c *gin.Context) {

	var input dto.EditUserComments
	//validate input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//UPDATE DB
	comment := domain.UserComments{Body: input.Body, CreatedOn: time.Now(), LastModified: time.Now()}
	if err := db.Model(&domain.UserComments{}).Where("comment_id = ?", input.CommentId).Omit("created_on").Updates(&comment).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, "SUCCESS")
	}
}