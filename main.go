package main

import (
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"sqlcTest/conn"
	"sqlcTest/modules"
)

type Credential struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main(){
	fmt.Println("postgres test crud")

	db := conn.GetConnection()

	db.AutoMigrate(&modules.Accounts{}, &modules.Books{}, &modules.UserPosts{}, &modules.UserComments{})

	r := gin.Default()
	r.POST("/account/login/", loginHandler)
	r.GET("/account/get/", auth, modules.GetAccount)
	r.GET("/account/get/:id", auth, modules.GetAccountId)
	r.POST("/account/add/", auth, modules.AddAccount)
	r.GET("/account/drop/:id", auth, modules.DelAccount)
	r.POST("/account/edit/:type", auth, modules.EditAccount)

	r.GET("/books/get/", auth, modules.GetBook)
	r.GET("/books/get/:id", auth, modules.GetBookId)
	r.POST("/books/add/", auth, modules.AddBook)
	r.GET("/books/drop/:id", auth, modules.DelBook)
	r.POST("/books/edit/:type", auth, modules.EditBook)

	r.GET("/comment/get/", auth, modules.GetComments)
	r.GET("/comment/get/:id", auth, modules.GetCommentById)
	r.GET("/comment/getByPost/:id", auth, modules.GetCommentByPostId)
	r.POST("/comment/add/", auth, modules.AddComments)
	r.GET("/comment/drop/:id", auth, modules.DelComment)
	r.POST("/comment/edit/:type", auth, modules.EditComment)

	r.GET("/post/get/", auth, modules.GetPost)
	r.GET("/post/get/:id", auth, modules.GetPostById)
	r.POST("/post/add/", auth, modules.AddPost)
	r.GET("/post/drop/:id", auth, modules.DelPost)
	r.POST("/post/edit/:type", auth, modules.EditPost)

	r.Run(":8080")
}

func loginHandler(c *gin.Context) {
	var user Credential
	err := c.Bind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "can't bind struct",
		})
	}
	if user.Username != "myname" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "wrong username or password",
		})
	} else {
		if user.Password != "myname123" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  http.StatusUnauthorized,
				"message": "wrong username or password",
			})
		}
	}
	sign := jwt.New(jwt.GetSigningMethod("HS256"))
	token, err := sign.SignedString([]byte("secret"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		c.Abort()
	}
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func auth(c *gin.Context) {
	tokenString := c.Request.Header.Get("Authorization")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != token.Method {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte("secret"), nil
	})

	if token != nil && err == nil {
		fmt.Println("token verified")
	} else {
		result := gin.H{
			"message": "not authorized",
			"error":   err.Error(),
		}
		c.JSON(http.StatusUnauthorized, result)
		c.Abort()
	}
}