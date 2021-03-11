package domain

import (
	"context"
	"sqlcTest/dto"
	"time"
)

type Accounts struct {
	UserId		int			`gorm:"primaryKey"`
	Username	string		`json:"username"`
	Password	string		`json:"password"`
	Email		string		`json:"email"`
	CreatedOn	time.Time	`json:"created_on"`
	LastLogin	time.Time	`json:"last_login"`
}

func (c Accounts) ToDto() dto.AccountResponse {
	return dto.AccountResponse{
		ID:    		c.UserId,
		Username:   c.Username,
		CreatedOn:  c.CreatedOn,
		LastLogin:  c.LastLogin,
	}
}

type AccountsRepository interface {
	FindByID(ctx context.Context, id int) (*Accounts, error)
	Fetch(ctx context.Context) ([]Accounts, error)
}

type AccountsService interface {
	Create(ctx context.Context, request dto.CreateAccounts) (*dto.AccountResponse, error)
	Edit(ctx context.Context, request dto.EditAccounts) (*dto.AccountResponse, error)
	Fetch(ctx context.Context) ([]dto.AccountResponse, error)
}
