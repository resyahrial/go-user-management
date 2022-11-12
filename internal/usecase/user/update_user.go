package usecase

import (
	"context"
	"strings"

	"github.com/resyahrial/go-user-management/internal/entities"
)

func (u *UserUsecaseImpl) Update(ctx context.Context, id string, input *entities.User) (user *entities.User, err error) {
	if strings.TrimSpace(input.Password) != "" {
		if input.Password, err = u.Hasher.HashPassword(input.Password); err != nil {
			return
		}
	}
	return u.UserRepo.Update(ctx, id, input)
}
