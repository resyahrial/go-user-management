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
	Role     *Role
}

type Role struct {
	Name        string
	Permissions []*Permission
}

type Permission struct {
	Resource string
	Action   string
	Type     string
}

func (p *Permission) IsGlobalPermission() bool {
	return p.Type == "GLOBAL"
}

type UserUsecase interface {
	Create(ctx context.Context, input *User) (user *User, err error)
	Update(ctx context.Context, id string, input *User) (user *User, err error)
	GetDetail(ctx context.Context, id string) (user *User, err error)
	GetList(ctx context.Context, params *PaginatedQueryParams) (users []*User, count int64, err error)
	Delete(ctx context.Context, id string) (err error)
}
