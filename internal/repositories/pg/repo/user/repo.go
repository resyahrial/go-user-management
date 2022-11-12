package repo

import (
	"context"

	"github.com/resyahrial/go-template/internal/entities"
	"github.com/resyahrial/go-template/internal/repositories/pg/models"
	"github.com/segmentio/ksuid"
	"gorm.io/gorm"
)

type UserRepoImpl struct {
	db *gorm.DB
}

func NewUserRepo(
	db *gorm.DB,
) *UserRepoImpl {
	return &UserRepoImpl{
		db,
	}
}

func (u *UserRepoImpl) Create(ctx context.Context, user *entities.User) (res *entities.User, err error) {
	var (
		userModel *models.User
	)

	if userModel, err = models.NewUserModel(user); err != nil {
		return
	}

	userModel.Id = ksuid.New().String()
	if err = u.db.WithContext(ctx).Create(userModel).Error; err != nil {
		return
	}

	return userModel.ConvertToEntity()
}
