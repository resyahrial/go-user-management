package response

import (
	"github.com/resyahrial/go-user-management/internal/entities"
)

type LoginResponse struct {
	AccessToken string `json:"access_token"`
}

func NewLoginResponse(token *entities.Token) (res *LoginResponse, err error) {
	res = &LoginResponse{
		AccessToken: token.Access,
	}
	return
}
