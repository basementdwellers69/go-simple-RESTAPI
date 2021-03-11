package dto

type CreateAccounts struct {
	Username	string		`json:"username" binding:"required"`
	Password	string		`json:"password" binding:"required"`
	Email		string		`json:"email" binding:"required"`
}
type EditAccounts struct {
	ID			int			`json:"user_id"`
	Username	string		`json:"username"`
	Password	string		`json:"password"`
	Email		string		`json:"email"`
}