package usecase

import (
	"context"
	"net/http"

	"github.com/resyahrial/go-user-management/internal/entities"
	"github.com/resyahrial/go-user-management/pkg/exception"
)

var (
	ErrInvalidToken = exception.NewAuthenticationException().SetModule(entities.AuthModule).SetMessage("token invalid, not contains user")
)

func (u *AuthUsecaseImpl) ValidateAccessToken(ctx context.Context, accessToken string) (user *entities.User, err error) {
	var (
		userId string
	)

	if userId, err = u.AccessTokenHandler.ParseToken(accessToken); err != nil {
		return
	}

	if user, err = u.UserRepo.GetById(ctx, userId); err != nil {
		if parsedErr, ok := err.(*exception.Base); ok && parsedErr.Code == http.StatusNotFound {
			err = ErrInvalidToken
		}
		return
	}

	return
}
