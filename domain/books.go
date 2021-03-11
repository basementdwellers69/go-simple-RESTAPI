package domain

import (
	"context"
	"sqlcTest/dto"
	"time"
)

type Books struct {
	BookId			int			`gorm:"primaryKey"`
	Title			string		`json:"title"`
	Synopsis		string		`json:"synopsis"`
	Tags			string		`json:"tags"`
	Author			string		`json:"author"`
	Price			int			`json:"price"`
	Status 			bool		`json:"status"`
	CreatedOn		time.Time	`json:"created_on"`
	LastModified	time.Time	`json:"last_modified"`
	Pages			int			`json:"pages"`
	Img				string		`json:"img"`
}

func (c Books) ToDto() dto.BooksResponse {
	return dto.BooksResponse{
		BookId:    		c.BookId,
		Tags:			c.Tags,
		Author:			c.Author,
		Price:			c.Price,
		Status: 		c.Status,
		LastModified:	"",
	}
}

type BooksRepository interface {
	FindByID(ctx context.Context, id int) (*Books, error)
	Fetch(ctx context.Context) ([]Books, error)
}

type BooksService interface {
	Create(ctx context.Context, request dto.CreateBooks) (*dto.BooksResponse, error)
	Edit(ctx context.Context, request dto.EditBooks) (*dto.BooksResponse, error)
	Fetch(ctx context.Context) ([]dto.BooksResponse, error)
}
