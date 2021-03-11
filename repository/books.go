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

func GetBook(c *gin.Context) {
	var book []domain.Books
	if err := db.Find(&book).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, book)
	}
}

func AddBook(c *gin.Context) {
	var input dto.CreateBooks
	//validate input
	if err := c.ShouldBindJSON(&input);err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//create Input
	book := domain.Books{
		Title:        input.Title,
		Synopsis:     input.Synopsis,
		Tags:         input.Tags,
		Author:       input.Author,
		Price:        input.Price,
		Status:       false,
		CreatedOn:    time.Now(),
		LastModified: time.Now(),
		Pages:        input.Pages,
		Img:          input.Img,
	}
	if err := db.Create(&book).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, book)
	}
}

func GetBookId(c *gin.Context) {
	id := c.Params.ByName("id")
	var book domain.Books
	if err := db.Where("book_id = ?", id).First(&book).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, book)
	}
}

func DelBook(c *gin.Context) {
	id := c.Params.ByName("id")
	var book domain.Books
	if err := db.Where("book_id = ?", id).Delete(&book).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, book)
	}
}
func EditBook(c *gin.Context) {
	//GET OMIT TYPES
	temp := c.Params.ByName("type")
	types := strings.Split(temp, ",")

	var input dto.EditBooks
	//validate input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//UPDATE DB
	book := domain.Books{Title: input.Title, Synopsis: input.Synopsis, Tags: input.Tags, Author: input.Author, Price: input.Price, Status: input.Status, CreatedOn: time.Now(), LastModified: time.Now()}
	if len(types) == 1 {
		if err := db.Model(&domain.Books{}).Where("book_id = ?", input.ID).Omit(types[0], "created_on").Updates(&book).Error; err != nil {
			c.AbortWithStatus(404)
			fmt.Println(err)
		} else {
			c.JSON(200, "SUCCESS")
		}
		return
	} else if len(types) == 2 {
		if err := db.Model(&domain.Books{}).Where("book_id = ?", input.ID).Omit(types[0], types[1], "created_on").Updates(&book).Error; err != nil {
			c.AbortWithStatus(404)
			fmt.Println(err)
		} else {
			c.JSON(200, "SUCCESS")
		}
		return
	} else if len(types) == 3 {
		if err := db.Model(&domain.Books{}).Where("book_id = ?", input.ID).Omit(types[0], types[1], types[2], "created_on").Updates(&book).Error; err != nil {
			c.AbortWithStatus(404)
			fmt.Println(err)
		} else {
			c.JSON(200, "SUCCESS")
		}
		return
	} else if len(types) == 4 {
		if err := db.Model(&domain.Books{}).Where("book_id = ?", input.ID).Omit(types[0], types[1],types[2], types[3], "created_on").Updates(&book).Error; err != nil {
			c.AbortWithStatus(404)
			fmt.Println(err)
		} else {
			c.JSON(200, "SUCCESS")
		}
		return
	} else if len(types) == 5 {
		if err := db.Model(&domain.Books{}).Where("book_id = ?", input.ID).Omit(types[0], types[1], types[2], types[3], types[4], "created_on").Updates(&book).Error; err != nil {
			c.AbortWithStatus(404)
			fmt.Println(err)
		} else {
			c.JSON(200, "SUCCESS")
		}
		return
	} else {
		c.JSON(200, "ERROR OCCURRED")
	}

}