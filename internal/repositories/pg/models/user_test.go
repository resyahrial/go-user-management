package models_test

import (
	"testing"

	"github.com/resyahrial/go-user-management/internal/entities"
	"github.com/resyahrial/go-user-management/internal/repositories/pg/models"
	"github.com/segmentio/ksuid"
	"github.com/stretchr/testify/assert"
)

func TestNewUserModel(t *testing.T) {
	userEntity := &entities.User{
		Name:     "user",
		Email:    "user@mail.com",
		Password: "anypassword",
		RoleName: "USER",
	}

	user, err := models.NewUserModel(userEntity)
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, userEntity.Name, user.Name)
	assert.Equal(t, userEntity.Email, user.Email)
	assert.Equal(t, userEntity.Password, user.Password)
	assert.Equal(t, userEntity.RoleName, user.RoleName)
}

func TestConvertToEntityUser(t *testing.T) {
	user := &models.User{
		CommonField: models.CommonField{
			ID: ksuid.New().String(),
		},
		Name:     "user",
		Email:    "user@mail.com",
		Password: "anypassword",
		Role: models.Role{
			Name: "ADMIN",
		},
	}

	userEntity, err := user.ConvertToEntity()
	assert.Nil(t, err)
	assert.NotNil(t, userEntity)
	assert.Equal(t, user.ID, userEntity.ID)
	assert.EqualValues(t, user.Name, userEntity.Name)
	assert.EqualValues(t, user.Email, userEntity.Email)
	assert.EqualValues(t, user.Password, userEntity.Password)
	assert.Equal(t, user.Role.Name, userEntity.Role.Name)
}
