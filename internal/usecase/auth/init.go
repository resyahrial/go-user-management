package usecase

import "github.com/resyahrial/go-user-management/internal/entities"

type AuthUsecaseImpl struct {
	UserRepo
	Hasher
	TokenHandler
}

func NewUserUsecase(
	userRepo UserRepo,
	hash Hasher,
	tokenHandler TokenHandler,
) entities.AuthUsecase {
	return &AuthUsecaseImpl{
		userRepo,
		hash,
		tokenHandler,
	}
}
