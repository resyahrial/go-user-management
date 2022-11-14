package usecase

import (
	"context"
	"strings"

	"github.com/resyahrial/go-user-management/internal/entities"
)

func (u *UserUsecaseImpl) Create(ctx context.Context, input *entities.User) (user *entities.User, err error) {
	if input.Password, err = u.Hasher.HashPassword(input.Password); err != nil {
		return
	}
	if strings.TrimSpace(input.RoleName) == "" {
		input.RoleName = entities.UserRole
	}
	return u.UserRepo.Create(ctx, input)
}
