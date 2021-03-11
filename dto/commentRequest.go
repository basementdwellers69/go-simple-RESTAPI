package dto

import "time"

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
