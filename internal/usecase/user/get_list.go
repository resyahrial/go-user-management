package usecase

import (
	"context"

	"github.com/resyahrial/go-user-management/internal/entities"
)

func (u *UserUsecaseImpl) GetList(ctx context.Context, params *entities.PaginatedQueryParams) (users []*entities.User, count int64, err error) {
	return u.UserRepo.GetList(ctx, params)
}
