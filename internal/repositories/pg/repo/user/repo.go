package repo

import (
	"context"
	"strings"

	"github.com/resyahrial/go-user-management/internal/entities"
	"github.com/resyahrial/go-user-management/internal/repositories/pg/models"
	"github.com/resyahrial/go-user-management/pkg/exception"
	"github.com/segmentio/ksuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var (
	ErrUserNotFound = exception.NewNotFoundException().SetModule(entities.UserModule).SetMessage("user not found")
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

	if err = u.db.WithContext(ctx).Model(&models.User{}).Where("id = ? AND is_deleted != true", id).First(&models.User{}).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			err = ErrUserNotFound
		}
		return
	}

	if userModel, err = models.NewUserModel(user); err != nil {
		return
	}

	if err = u.db.WithContext(ctx).Model(&userModel).Clauses(clause.Returning{}).Where("id = ?", id).Updates(userModel).Error; err != nil {
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

func (u *UserRepoImpl) GetById(ctx context.Context, id string) (res *entities.User, err error) {
	var (
		userModel *models.User
	)

	if err = u.db.WithContext(ctx).Model(&models.User{}).Where("id = ? AND is_deleted != true", id).First(&userModel).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			err = ErrUserNotFound
		}
		return
	}

	return userModel.ConvertToEntity()
}

func (u *UserRepoImpl) GetList(ctx context.Context, params *entities.PaginatedQueryParams) (users []*entities.User, count int64, err error) {
	var (
		userModels []models.User
	)

	result := u.db.WithContext(ctx).Where("is_deleted != true")
	if err = result.Model(&models.User{}).Count(&count).Error; err != nil {
		return
	}

	if err = result.Limit(params.Limit).Offset(params.Limit * params.Page).Find(&userModels).Error; err != nil {
		return
	}

	users = make([]*entities.User, 0)
	for _, userModel := range userModels {
		if user, errConvert := userModel.ConvertToEntity(); errConvert != nil {
			err = errConvert
			return nil, 0, err
		} else {
			users = append(users, user)
		}
	}

	return
}

func (u *UserRepoImpl) Delete(ctx context.Context, id string) (err error) {
	result := u.db.WithContext(ctx).Model(&models.User{}).Where("id = ? AND is_deleted != true", id)
	if err = result.First(&models.User{}).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			err = ErrUserNotFound
		}
		return
	}

	if err = result.Clauses(clause.Returning{}).Updates(&models.User{
		CommonField: models.CommonField{
			IsDeleted: true,
		},
	}).Error; err != nil {
		return
	}

	return
}
