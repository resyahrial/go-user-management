package usecase

import (
	"context"

	"github.com/resyahrial/go-user-management/internal/entities"
)

//go:generate mockgen -destination=mocks/mock.go -source=adapter.go UserRepo
type UserRepo interface {
	GetByEmail(ctx context.Context, email string) (res *entities.User, err error)
	GetById(ctx context.Context, id string) (res *entities.User, err error)
	GetByIdWithPermission(ctx context.Context, id string) (res *entities.User, err error)
}

//go:generate mockgen -destination=mocks/mock.go -source=adapter.go Hasher
type Hasher interface {
	CheckPasswordHash(password, hash string) bool
}

//go:generate mockgen -destination=mocks/mock.go -source=adapter.go TokenHandler
type TokenHandler interface {
	SignToken(claims map[string]interface{}) (tokenString string, err error)
	ParseToken(tokenString string) (id string, err error)
}
