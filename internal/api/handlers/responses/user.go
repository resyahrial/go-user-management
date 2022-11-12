package response

import (
	"github.com/mitchellh/mapstructure"
	"github.com/resyahrial/go-user-management/internal/entities"
)

type CreateUserResponse struct {
	Name  string
	Email string
}

func NewCreateUserResponse(user *entities.User) (res *CreateUserResponse, err error) {
	if err = mapstructure.Decode(user, &res); err != nil {
		return
	}
	return
}
