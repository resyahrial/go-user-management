//go:build wireinject
// +build wireinject

package factory

import (
	"github.com/google/wire"
	"github.com/resyahrial/go-user-management/internal/entities"
	repo "github.com/resyahrial/go-user-management/internal/repositories/pg/repo/user"
	usecase "github.com/resyahrial/go-user-management/internal/usecase/user"
	"gorm.io/gorm"
)

func InitUserUsecase(db *gorm.DB) entities.UserUsecase {
	wire.Build(
		usecase.NewUserUsecase,
		userRepoAdapterSet,
	)
	return nil
}

var userRepoAdapterSet = wire.NewSet(
	repo.NewUserRepo,
	wire.Bind(new(usecase.UserRepo), new(*repo.UserRepoImpl)),
)
