//go:build wireinject
// +build wireinject

package factory

import (
	"github.com/google/wire"
	"gorm.io/gorm"
	"time"

	"github.com/resyahrial/go-user-management/internal/entities"
	user_repo "github.com/resyahrial/go-user-management/internal/repositories/pg/repo/user"
	auth_usecase "github.com/resyahrial/go-user-management/internal/usecase/auth"
	user_usecase "github.com/resyahrial/go-user-management/internal/usecase/user"
	"github.com/resyahrial/go-user-management/pkg/authorization"
	"github.com/resyahrial/go-user-management/pkg/hasher"
)

func InitUserUsecase(db *gorm.DB, hasherCost int) entities.UserUsecase {
	wire.Build(
		user_usecase.NewUserUsecase,
		userRepoUserUsecaseAdapterSet,
		hasherUserUsecaseAdapterSet,
	)
	return nil
}

var userRepoUserUsecaseAdapterSet = wire.NewSet(
	user_repo.NewUserRepo,
	wire.Bind(new(user_usecase.UserRepo), new(*user_repo.UserRepoImpl)),
)

var hasherUserUsecaseAdapterSet = wire.NewSet(
	hasher.New,
	wire.Bind(new(user_usecase.Hasher), new(*hasher.Hasher)),
)

func InitAuthUsecase(db *gorm.DB, hasherCost int, tokenDuration time.Duration, secretKey string) entities.AuthUsecase {
	wire.Build(
		auth_usecase.NewAuthUsecase,
		userRepoAuthUsecaseAdapterSet,
		hasherAuthUsecaseAdapterSet,
		tokenHandlerAdapterSet,
	)
	return nil
}

var userRepoAuthUsecaseAdapterSet = wire.NewSet(
	user_repo.NewUserRepo,
	wire.Bind(new(auth_usecase.UserRepo), new(*user_repo.UserRepoImpl)),
)

var hasherAuthUsecaseAdapterSet = wire.NewSet(
	hasher.New,
	wire.Bind(new(auth_usecase.Hasher), new(*hasher.Hasher)),
)

var tokenHandlerAdapterSet = wire.NewSet(
	authorization.NewJwtAuthorization,
	wire.Bind(new(auth_usecase.TokenHandler), new(*authorization.JwtAuthorization)),
)
