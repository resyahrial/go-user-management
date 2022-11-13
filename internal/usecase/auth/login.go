package usecase

import (
	"context"

	"github.com/resyahrial/go-user-management/internal/entities"
	"github.com/resyahrial/go-user-management/pkg/exception"
)

var (
	ErrInvalidLoginIput = exception.NewAuthenticationException().SetModule(entities.AuthModule).SetMessage("email / password is invalid")
)

func (u *AuthUsecaseImpl) Login(ctx context.Context, input *entities.Login) (token *entities.Token, err error) {
	var (
		user        *entities.User
		accessToken string
	)

	if user, err = u.UserRepo.GetByEmail(ctx, input.Email); err != nil {
		return
	}

	if ok := u.Hasher.CheckPasswordHash(input.Password, user.Password); !ok {
		err = ErrInvalidLoginIput
		return
	}

	if accessToken, err = u.TokenHandler.SignToken(map[string]interface{}{
		"id": user.ID,
	}); err != nil {
		return
	} else {
		token = &entities.Token{
			Access: accessToken,
		}
	}

	return
}
