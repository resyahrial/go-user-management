package models

import (
	"github.com/mitchellh/mapstructure"
	"github.com/resyahrial/go-user-management/internal/entities"
)

type User struct {
	CommonField
	Name     string
	Email    string
	Password string
	RoleName string
	Role     Role `gorm:"foreignKey:RoleName;references:Name"`
}

func NewUserModel(userEntity *entities.User) (user *User, err error) {
	if err = mapstructure.Decode(userEntity, &user); err != nil {
		return
	}
	return
}

func (u *User) ConvertToEntity() (userEntity *entities.User, err error) {
	if err = mapstructure.Decode(u, &userEntity); err != nil {
		return
	}
	userEntity.ID = u.ID
	if u.Role.Name == "" {
		userEntity.Role = nil
	}
	return
}
