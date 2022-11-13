package models_test

import (
	"testing"

	"github.com/resyahrial/go-user-management/internal/repositories/pg/models"
	"github.com/segmentio/ksuid"
	"github.com/stretchr/testify/assert"
)

func TestConvertToPermissionEntity(t *testing.T) {
	permission := &models.Permission{
		ID:       ksuid.New().String(),
		Resource: "users",
		Action:   "WRITE",
		Type:     "GLOBAL",
	}

	permissionEntity, err := permission.ConvertToEntity()
	assert.Nil(t, err)
	assert.NotNil(t, permissionEntity)
	assert.Equal(t, permission.Resource, permissionEntity.Resource)
	assert.Equal(t, permission.Action, permissionEntity.Action)
	assert.Equal(t, permission.Type, permissionEntity.Type)
}
