package dto

import (
	"github.com/lib/pq"
	"time"
)

type PostsResponse struct {
	PostId			int				`json:"post_id"`
	Username 		string			`json:"username"`
	Body			string			`json:"body"`
	Likes			int				`json:"likes"`
	Img 			pq.StringArray	`json:"img" gorm:"type:string[]"`
	CreatedOn		time.Time		`json:"created_on"`
	LastModified	time.Time		`json:"last_modified"`
}
