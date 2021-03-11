package domain

import (
	"context"
	"github.com/lib/pq"
	"sqlcTest/dto"
	"time"
)

type UserPosts struct {
	PostId			int				`gorm:"primaryKey"`
	UserId 			int				`json:"user_id"`
	Body			string			`json:"body"`
	Likes			int				`json:"likes"`
	CreatedOn		time.Time		`json:"created_on"`
	LastModified	time.Time		`json:"last_modified"`
	Img				pq.StringArray	`json:"img" gorm:"type:string[]"`
}

func (c UserPosts) ToDto() dto.PostsResponse {
	return dto.PostsResponse{
		Username: 		"",
		Body:			c.Body,
		Likes:			c.Likes,
		CreatedOn:		c.CreatedOn,
		LastModified:	c.LastModified,
	}
}

type PostsRepository interface {
	FindByID(ctx context.Context, id int) (*UserPosts, error)
	Fetch(ctx context.Context) ([]UserPosts, error)
}

type PostsService interface {
	Create(ctx context.Context, request dto.CreateUserPosts) (*dto.PostsResponse, error)
	Edit(ctx context.Context, request dto.EditUserPosts) (*dto.PostsResponse, error)
	Fetch(ctx context.Context) ([]dto.PostsResponse, error)
}
