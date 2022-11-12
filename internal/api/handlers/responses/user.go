package response

import (
	"github.com/mitchellh/mapstructure"
	"github.com/resyahrial/go-user-management/internal/entities"
)

type CreateUserResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func NewCreateUserResponse(user *entities.User) (res *CreateUserResponse, err error) {
	if err = mapstructure.Decode(user, &res); err != nil {
		return
	}
	return
}
