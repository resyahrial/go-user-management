package usecase

import (
	"context"

	"github.com/resyahrial/go-user-management/internal/entities"
)

//go:generate mockgen -destination=mocks/mock.go -source=adapter.go UserRepo
type UserRepo interface {
	Create(ctx context.Context, user *entities.User) (res *entities.User, err error)
	Update(ctx context.Context, id string, user *entities.User) (res *entities.User, err error)
	GetById(ctx context.Context, id string) (res *entities.User, err error)
	GetList(ctx context.Context, params *entities.PaginatedQueryParams) (users []*entities.User, count int64, err error)
}

//go:generate mockgen -destination=mocks/mock.go -source=adapter.go Hasher
type Hasher interface {
	HashPassword(password string) (string, error)
}
