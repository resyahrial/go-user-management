package usecase

import "github.com/resyahrial/go-user-management/internal/entities"

type AuthUsecaseImpl struct {
	UserRepo
	Hasher
	AccessTokenHandler TokenHandler
}

func NewAuthUsecase(
	userRepo UserRepo,
	hash Hasher,
	accessTokenHandler TokenHandler,
) entities.AuthUsecase {
	return &AuthUsecaseImpl{
		userRepo,
		hash,
		accessTokenHandler,
	}
}
