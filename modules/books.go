package modules

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

type Books struct {
	BookId			int			`gorm:"primaryKey"`
	Title			string		`json:"username"`
	Synopsis		string		`json:"password"`
	Tags			string		`json:"tags"`
	Author			string		`json:"author"`
	Price			int			`json:"price"`
	Status 			bool		`json:"status"`
	CreatedOn		time.Time	`json:"created_on"`
	LastModified	time.Time	`json:"last_modified"`
}

type CreateBooks struct {
	Title			string		`json:"username" binding:"required"`
	Synopsis		string		`json:"password" binding:"required"`
	Tags			string		`json:"tags" binding:"required"`
	Author			string		`json:"author" binding:"required"`
	Price			int			`json:"price" binding:"required"`
}
type EditBooks struct {
	ID				int			`json:"book_id"`
	Title			string		`json:"username"`
	Synopsis		string		`json:"password"`
	Tags			string		`json:"tags"`
	Author			string		`json:"author"`
	Price			int			`json:"price"`
	Status 			bool		`json:"status"`
	LastModified	time.Time	`json:"last_modified"`
}

func GetBook(c *gin.Context) {
	var book []Books
	if err := db.Find(&book).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, book)
	}
}

func AddBook(c *gin.Context) {

	var input CreateBooks
	//validate input
	if err := c.ShouldBindJSON(&input);err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//create Input
	book := Books{Title: input.Title, Synopsis: input.Synopsis, Tags: input.Tags, Author: input.Author, Price: input.Price, Status: false, CreatedOn: time.Now(), LastModified: time.Now()}
	if err := db.Create(&book).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, book)
	}
}

func GetBookId(c *gin.Context) {
	id := c.Params.ByName("id")
	var book Books
	if err := db.Where("book_id = ?", id).First(&book).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, book)
	}
}

func DelBook(c *gin.Context) {
	id := c.Params.ByName("id")
	var book Books
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

	var input EditBooks
	//validate input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//UPDATE DB
	book := Books{Title: input.Title, Synopsis: input.Synopsis, Tags: input.Tags, Author: input.Author, Price: input.Price, Status: input.Status, CreatedOn: time.Now(), LastModified: time.Now()}
	if len(types) == 1 {
		if err := db.Model(&Books{}).Where("book_id = ?", input.ID).Omit(types[0], "created_on").Updates(&book).Error; err != nil {
			c.AbortWithStatus(404)
			fmt.Println(err)
		} else {
			c.JSON(200, "SUCCESS")
		}
		return
	} else if len(types) == 2 {
		if err := db.Model(&Books{}).Where("book_id = ?", input.ID).Omit(types[0], types[1], "created_on").Updates(&book).Error; err != nil {
			c.AbortWithStatus(404)
			fmt.Println(err)
		} else {
			c.JSON(200, "SUCCESS")
		}
		return
	} else if len(types) == 3 {
		if err := db.Model(&Books{}).Where("book_id = ?", input.ID).Omit(types[0], types[1], types[2], "created_on").Updates(&book).Error; err != nil {
			c.AbortWithStatus(404)
			fmt.Println(err)
		} else {
			c.JSON(200, "SUCCESS")
		}
		return
	} else if len(types) == 4 {
		if err := db.Model(&Books{}).Where("book_id = ?", input.ID).Omit(types[0], types[1],types[2], types[3], "created_on").Updates(&book).Error; err != nil {
			c.AbortWithStatus(404)
			fmt.Println(err)
		} else {
			c.JSON(200, "SUCCESS")
		}
		return
	} else if len(types) == 5 {
		if err := db.Model(&Books{}).Where("book_id = ?", input.ID).Omit(types[0], types[1], types[2], types[3], types[4], "created_on").Updates(&book).Error; err != nil {
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