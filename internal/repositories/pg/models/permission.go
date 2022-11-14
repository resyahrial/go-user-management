package models

import (
	"github.com/mitchellh/mapstructure"
	"github.com/resyahrial/go-user-management/internal/entities"
)

type Permission struct {
	ID       string `gorm:"primaryKey"`
	Resource string
	Action   string
	Type     string
}

func (p *Permission) ConvertToEntity() (permissionEntity *entities.Permission, err error) {
	if err = mapstructure.Decode(p, &permissionEntity); err != nil {
		return
	}
	return
}
