package dto

type BooksResponse struct {
	BookId			int			`json:"book_id"`
	Title			string		`json:"title"`
	Synopsis		string		`json:"synopsis"`
	Tags			string		`json:"tags"`
	Author			string		`json:"author"`
	Price			int			`json:"price"`
	Status 			bool		`json:"status"`
	CreatedOn		string		`json:"created_on"`
	LastModified	string		`json:"last_modified"`
	Pages			int			`json:"pages"`
	Img				string		`json:"img"`
}
