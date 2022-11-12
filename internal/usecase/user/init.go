package usecase

import "github.com/resyahrial/go-template/internal/entities"

type UserUsecaseImpl struct {
	UserRepo
}

func NewUserUsecase(
	userRepo UserRepo,
) entities.UserUsecase {
	return &UserUsecaseImpl{
		UserRepo: userRepo,
	}
}
