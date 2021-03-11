package dto

import (
	"github.com/lib/pq"
	"time"
)

type CreateUserPosts struct {
	UserId 			int				`json:"user_id" binding:"required"`
	Body			string			`json:"body"  binding:"required"`
	Img				pq.StringArray	`json:"img" gorm:"type:string[]"`
}

type EditUserPosts struct {
	PostId			string			`json:"post_id"`
	Body			string			`json:"body"`
	Img 			pq.StringArray 	`json:"img" gorm:"type:string[]"`
	LastModified	time.Time		`json:"last_modified"`
}
