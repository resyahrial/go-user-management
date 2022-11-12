package usecase

import "github.com/resyahrial/go-user-management/internal/entities"

type UserUsecaseImpl struct {
	UserRepo
	Hasher
}

func NewUserUsecase(
	userRepo UserRepo,
	hash Hasher,
) entities.UserUsecase {
	return &UserUsecaseImpl{
		userRepo,
		hash,
	}
}
