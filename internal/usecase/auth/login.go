package usecase

import (
	"context"

	"github.com/resyahrial/go-user-management/internal/entities"
)

func (u *AuthUsecaseImpl) Login(ctx context.Context, input *entities.Login) (token *entities.Token, err error) {
	return
}
