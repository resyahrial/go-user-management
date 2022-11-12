//go:build wireinject
// +build wireinject

package factory

import (
	"github.com/google/wire"
	"github.com/resyahrial/go-template/internal/entities"
	repo "github.com/resyahrial/go-template/internal/repositories/pg/repo/user"
	usecase "github.com/resyahrial/go-template/internal/usecase/user"
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
