package dto

import "time"

type CommentsResponse struct {
	CommentId		int			`gorm:"primaryKey"`
	PostId			int			`json:"post_id"`
	Username 		string		`json:"username"`
	Body			string		`json:"body"`
	CreatedOn		time.Time	`json:"created_on"`
	LastModified	time.Time	`json:"last_modified"`
}
