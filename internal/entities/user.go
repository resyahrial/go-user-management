package entities

import (
	"context"
)

const (
	UserModule = "USER"
)

type User struct {
	Id       string
	Name     string
	Email    string
	Password string
}

type UserUsecase interface {
	CreateUser(ctx context.Context, input *User) (user *User, err error)
}
