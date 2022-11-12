//go:build wireinject
// +build wireinject

package factory

import (
	"github.com/google/wire"
	"github.com/resyahrial/go-user-management/internal/entities"
	repo "github.com/resyahrial/go-user-management/internal/repositories/pg/repo/user"
	usecase "github.com/resyahrial/go-user-management/internal/usecase/user"
	"github.com/resyahrial/go-user-management/pkg/hasher"
	"gorm.io/gorm"
)

func InitUserUsecase(db *gorm.DB, hasherCost int) entities.UserUsecase {
	wire.Build(
		usecase.NewUserUsecase,
		userRepoAdapterSet,
		hasherAdapterSet,
	)
	return nil
}

var userRepoAdapterSet = wire.NewSet(
	repo.NewUserRepo,
	wire.Bind(new(usecase.UserRepo), new(*repo.UserRepoImpl)),
)

var hasherAdapterSet = wire.NewSet(
	hasher.New,
	wire.Bind(new(usecase.Hasher), new(*hasher.Hasher)),
)
