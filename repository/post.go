package repository

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sqlcTest/domain"
	"sqlcTest/dto"
	"strings"
	"time"
)

func GetPost(c *gin.Context) {
	var post []domain.UserPosts
	var res []dto.PostsResponse
	if err := db.Model(&post).Select("post_id, accounts.username, body, likes, user_posts.img, user_posts.created_on, user_posts.last_modified").Joins("INNER JOIN accounts ON user_posts.user_id = accounts.user_id").Find(&res).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, res)
	}
}

func AddPost(c *gin.Context) {

	var input dto.CreateUserPosts
	//validate input
	if err := c.ShouldBindJSON(&input);err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//create Input
	post:= domain.UserPosts{UserId:input.UserId, Body: input.Body, Likes: 0, CreatedOn: time.Now(), LastModified: time.Now(), Img: input.Img}
	if err := db.Create(&post).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, post)
	}
}

func GetPostById(c *gin.Context) {
	id := c.Params.ByName("id")
	var post domain.UserPosts
	var res []dto.PostsResponse
	if err := db.Model(&post).Select("post_id, accounts.username, body, likes, user_posts.img, user_posts.created_on, user_posts.last_modified").Joins("INNER JOIN accounts ON user_posts.user_id = accounts.user_id").Where("post_id = ?", id).First(&res).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, res)
	}
}

func DelPost(c *gin.Context) {
	id := c.Params.ByName("id")
	var post domain.UserPosts
	if err := db.Where("post_id = ?", id).Delete(&post).Error; err != nil {
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

	var input dto.EditUserPosts
	//validate input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//UPDATE DB
	post := domain.UserPosts{Body: input.Body, Likes: 0, Img: input.Img, CreatedOn: time.Now(), LastModified: time.Now()}
	if len(types) == 1 {
		if err := db.Model(&domain.UserPosts{}).Where("post_id = ?", input.PostId).Omit(types[0], "created_on").Updates(&post).Error; err != nil {
			c.AbortWithStatus(404)
			fmt.Println(err)
		} else {
			c.JSON(200, "SUCCESS")
		}
	} else if len(types) == 2 {
			if err := db.Model(&domain.UserPosts{}).Where("post_id = ?", input.PostId).Omit(types[0], types[1], "created_on").Updates(&post).Error; err != nil {
				c.AbortWithStatus(404)
				fmt.Println(err)
			} else {
				c.JSON(200, "SUCCESS")
			}
	} else {
			c.JSON(200, "ERROR")
		}
	}