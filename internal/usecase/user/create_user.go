package usecase

import (
	"context"

	"github.com/resyahrial/go-template/internal/entities"
)

func (u *UserUsecaseImpl) CreateUser(ctx context.Context, input *entities.User) (user *entities.User, err error) {
	return u.UserRepo.Create(ctx, input)
}
