package dto

import "time"

type AccountResponse struct {
	ID			int			`json:"user_id"`
	Username	string		`json:"username"`
	CreatedOn	time.Time	`json:"created_on"`
	LastLogin	time.Time	`json:"last_login"`
}
