package request

import (
	"github.com/mitchellh/mapstructure"
	"github.com/resyahrial/go-user-management/internal/entities"
	"github.com/resyahrial/go-user-management/pkg/exception"
	"github.com/resyahrial/go-user-management/pkg/validator"
)

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,alphanum,min=8"`
}

func (r *LoginRequest) CastToLoginEntity() (login *entities.Login, err error) {
	if mapErr := validator.Validate(r); len(mapErr) > 0 {
		err = exception.NewBadRequestException().SetModule(entities.AuthModule).SetCollectionMessage(mapErr)
		return
	}
	if err = mapstructure.Decode(r, &login); err != nil {
		return
	}
	return
}
