package v1

import (
	"time"

	"github.com/resyahrial/go-user-management/config"
	"github.com/resyahrial/go-user-management/internal/entities"
	"github.com/resyahrial/go-user-management/internal/factory"
	"gorm.io/gorm"
)

type Handler struct {
	userUsecase entities.UserUsecase
	authUsecase entities.AuthUsecase
}

func NewHandler(cfg config.Config, db *gorm.DB) *Handler {
	timeDuration := 10 * time.Second
	secret := "secret"
	return &Handler{
		userUsecase: factory.InitUserUsecase(db, cfg.Hasher.Cost),
		authUsecase: factory.InitAuthUsecase(db, cfg.Hasher.Cost, timeDuration, secret),
	}
}
