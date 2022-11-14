package usecase

import (
	"context"

	"github.com/resyahrial/go-user-management/internal/entities"
	"github.com/resyahrial/go-user-management/pkg/exception"
)

var (
	ErrUnauthenticated = exception.NewAuthenticationException().SetModule(entities.AuthModule).SetMessage("user didn't have permission on this resource")
)

func (u *AuthUsecaseImpl) ValidateUserAccess(ctx context.Context, authentication *entities.Authentication) (err error) {
	var (
		user *entities.User
	)

	if user, err = u.UserRepo.GetByIdWithPermission(ctx, authentication.CurrentUserID); err != nil {
		return
	}

	if !authentication.IsPermissionValid(user) {
		err = ErrUnauthenticated
		return
	}

	return
}
