package domain

import (
	"context"
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

func (c UserComments) ToDto() dto.CommentsResponse {
	return dto.CommentsResponse{
		PostId:    		c.PostId,
		Username:  		"",
		Body:			c.Body,
		CreatedOn:		c.CreatedOn,
		LastModified:	c.LastModified,
	}
}

type CommentsRepository interface {
	FindByPostID(ctx context.Context, id int) (*UserComments, error)
	Fetch(ctx context.Context) ([]UserComments, error)
}

type CommentsService interface {
	Create(ctx context.Context, request dto.CreateUserComments) (*dto.CommentsResponse, error)
	Edit(ctx context.Context, request dto.EditUserComments) (*dto.CommentsResponse, error)
	Fetch(ctx context.Context) ([]dto.CommentsResponse, error)
}
