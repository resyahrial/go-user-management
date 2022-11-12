package usecase

import (
	"context"

	"github.com/resyahrial/go-user-management/internal/entities"
)

//go:generate mockgen -destination=mocks/mock.go -source=adapter.go UserRepo
type UserRepo interface {
	Create(ctx context.Context, user *entities.User) (res *entities.User, err error)
}

//go:generate mockgen -destination=mocks/mock.go -source=adapter.go Hasher
type Hasher interface {
	HashPassword(password string) (string, error)
}