package usecase

import (
	"context"

	"github.com/resyahrial/go-template/internal/entities"
)

//go:generate mockgen -destination=mocks/mock.go -source=adapter.go UserRepo
type UserRepo interface {
	Create(ctx context.Context, user *entities.User) (res *entities.User, err error)
}
