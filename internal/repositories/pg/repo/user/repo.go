package repo

import (
	"context"
	"strings"

	"github.com/resyahrial/go-user-management/internal/entities"
	"github.com/resyahrial/go-user-management/internal/repositories/pg/models"
	"github.com/resyahrial/go-user-management/pkg/exception"
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
	mapError := make(map[string][]string)

	if userModel, err = models.NewUserModel(user); err != nil {
		return
	}

	userModel.ID = ksuid.New().String()
	if err = u.db.WithContext(ctx).Create(userModel).Error; err != nil {
		if strings.Contains(err.Error(), `duplicate key value violates unique constraint "users_email_key"`) {
			mapError["email"] = []string{
				"Email must be unique",
			}
		}
		if len(mapError) > 0 {
			err = exception.NewBadRequestException().SetModule(entities.UserModule).SetCollectionMessage(mapError)
		}
		return
	}

	return userModel.ConvertToEntity()
}

func (u *UserRepoImpl) Update(ctx context.Context, id string, user *entities.User) (res *entities.User, err error) {
	var (
		userModel *models.User
	)
	mapError := make(map[string][]string)

	if userModel, err = models.NewUserModel(user); err != nil {
		return
	}

	userModel.ID = id
	if err = u.db.WithContext(ctx).Model(user).Updates(&userModel).Error; err != nil {
		if strings.Contains(err.Error(), `duplicate key value violates unique constraint "users_email_key"`) {
			mapError["email"] = []string{
				"Email must be unique",
			}
		}
		if len(mapError) > 0 {
			err = exception.NewBadRequestException().SetModule(entities.UserModule).SetCollectionMessage(mapError)
		}
		return
	}

	return userModel.ConvertToEntity()
}
