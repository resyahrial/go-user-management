package response

import (
	"github.com/mitchellh/mapstructure"
	"github.com/resyahrial/go-user-management/internal/entities"
)

type UserResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func NewUserResponse(user *entities.User) (res *UserResponse, err error) {
	if err = mapstructure.Decode(user, &res); err != nil {
		return
	}
	return
}

func NewListUserResponse(users []*entities.User) (res []*UserResponse, err error) {
	res = make([]*UserResponse, 0)
	for _, user := range users {
		var userRes *UserResponse
		if userRes, err = NewUserResponse(user); err != nil {
			return nil, err
		}
		res = append(res, userRes)
	}
	return
}
