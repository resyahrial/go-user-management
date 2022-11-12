package v1

import (
	"github.com/resyahrial/go-user-management/internal/entities"
	"github.com/resyahrial/go-user-management/internal/factory"
	"gorm.io/gorm"
)

type Handler struct {
	userUsecase entities.UserUsecase
}

func NewHandler(db *gorm.DB, hashCost int) *Handler {
	return &Handler{
		userUsecase: factory.InitUserUsecase(db, hashCost),
	}
}
