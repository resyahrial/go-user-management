package entities

import "context"

const (
	AuthModule = "AUTH"
)

type Login struct {
	Email    string
	Password string
}

type Token struct {
	Access string
}

type AuthUsecase interface {
	Login(ctx context.Context, input *Login) (token *Token, err error)
	ValidateAccessToken(ctx context.Context, accessToken string) (user *User, err error)
}
