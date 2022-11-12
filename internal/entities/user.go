package entities

import (
	"context"
)

const (
	UserModule = "USER"
)

type User struct {
	ID       string
	Name     string
	Email    string
	Password string
}

type UserUsecase interface {
	Create(ctx context.Context, input *User) (user *User, err error)
	Update(ctx context.Context, id string, input *User) (user *User, err error)
	GetDetail(ctx context.Context, id string) (user *User, err error)
	GetList(ctx context.Context, param *PaginatedQueryParams) (users []*User, count int64, err error)
}
