package dto

import "time"

type CreateBooks struct {
	Title			string		`json:"title" binding:"required"`
	Synopsis		string		`json:"synopsis" binding:"required"`
	Tags			string		`json:"tags" binding:"required"`
	Author			string		`json:"author" binding:"required"`
	Price			int			`json:"price" binding:"required"`
	Pages			int			`json:"pages" binding:"required"`
	Img				string		`json:"img"`
}
type EditBooks struct {
	ID				int			`json:"book_id"`
	Title			string		`json:"title"`
	Synopsis		string		`json:"synopsis"`
	Tags			string		`json:"tags"`
	Author			string		`json:"author"`
	Price			int			`json:"price"`
	Status 			bool		`json:"status"`
	LastModified	time.Time	`json:"last_modified"`
	Pages			int			`json:"pages" binding:"required"`
	Img				string		`json:"img"`
}
