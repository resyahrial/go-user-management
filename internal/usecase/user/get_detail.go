package usecase

import (
	"context"

	"github.com/resyahrial/go-user-management/internal/entities"
)

func (u *UserUsecaseImpl) GetDetail(ctx context.Context, id string) (user *entities.User, err error) {
	return u.UserRepo.GetById(ctx, id)
}
