package usecase

import (
	"context"
)

func (u *UserUsecaseImpl) Delete(ctx context.Context, id string) (err error) {
	return u.UserRepo.Delete(ctx, id)
}
