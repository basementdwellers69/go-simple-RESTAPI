package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"sqlcTest/config"
	"sqlcTest/domain"
	"sqlcTest/handler"
	"sqlcTest/repository"
	"sqlcTest/service"
)

func Start() {
	fmt.Println("postgres test crud")

	db := config.GetConnection()

	db.AutoMigrate(&domain.Accounts{}, &domain.Books{}, &domain.UserPosts{}, &domain.UserComments{})

	r := gin.Default()
	r.POST("/account/login/", handler.LoginHandler)
	r.GET("/account/get/", service.Auth, repository.GetAccount)
	r.GET("/account/get/:id", service.Auth, repository.GetAccountId)
	r.POST("/account/add/", service.Auth, repository.AddAccount)
	r.GET("/account/drop/:id", service.Auth, repository.DelAccount)
	r.POST("/account/edit/:type", service.Auth, repository.EditAccount)

	r.GET("/books/get/", service.Auth, repository.GetBook)
	r.GET("/books/get/:id", service.Auth, repository.GetBookId)
	r.POST("/books/add/", service.Auth, repository.AddBook)
	r.GET("/books/drop/:id", service.Auth, repository.DelBook)
	r.POST("/books/edit/:type", service.Auth, repository.EditBook)

	r.GET("/comment/get/", service.Auth, repository.GetComments)
	r.GET("/comment/get/:id", service.Auth, repository.GetCommentById)
	r.GET("/comment/getByPost/:id", service.Auth, repository.GetCommentByPostId)
	r.POST("/comment/add/", service.Auth, repository.AddComments)
	r.GET("/comment/drop/:id", service.Auth, repository.DelComment)
	r.POST("/comment/edit/:type", service.Auth, repository.EditComment)

	r.GET("/post/get/", service.Auth, repository.GetPost)
	r.GET("/post/get/:id", service.Auth, repository.GetPostById)
	r.POST("/post/add/", service.Auth, repository.AddPost)
	r.GET("/post/drop/:id", service.Auth, repository.DelPost)
	r.POST("/post/edit/:type", service.Auth, repository.EditPost)

	r.Run(":8080")
}
