package request

import (
	"github.com/mitchellh/mapstructure"
	"github.com/resyahrial/go-user-management/internal/entities"
	"github.com/resyahrial/go-user-management/pkg/exception"
	"github.com/resyahrial/go-user-management/pkg/validator"
)

type CreateUserRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,alphanum,min=8"`
}

func (r *CreateUserRequest) CastToUserEntity() (user *entities.User, err error) {
	if mapErr := validator.Validate(r); len(mapErr) > 0 {
		err = exception.NewBadRequestException().SetModule(entities.UserModule).SetCollectionMessage(mapErr)
		return
	}
	if err = mapstructure.Decode(r, &user); err != nil {
		return
	}
	return
}

type UpdateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email" validate:"omitempty,email"`
	Password string `json:"password" validate:"omitempty,alphanum,min=8"`
}

func (r *UpdateUserRequest) CastToUserEntity() (user *entities.User, err error) {
	if mapErr := validator.Validate(r); len(mapErr) > 0 {
		err = exception.NewBadRequestException().SetModule(entities.UserModule).SetCollectionMessage(mapErr)
		return
	}
	if err = mapstructure.Decode(r, &user); err != nil {
		return
	}
	return
}
