package seeds

import (
	"github.com/resyahrial/go-user-management/config"
	"github.com/resyahrial/go-user-management/internal/repositories/pg/models"
	"github.com/resyahrial/go-user-management/pkg/hasher"
	"gorm.io/gorm"
)

func seedAdmin(db *gorm.DB, cfg config.Config) (err error) {
	if !db.Migrator().HasTable(&models.User{}) {
		return
	}

	admin := &models.User{
		Name:     "Admin",
		Email:    cfg.Seed.Admin.Email,
		RoleName: "ADMIN",
	}

	// check if any admins created, continue if not any admins created
	var count int64
	if err = db.Model(&models.User{}).Count(&count).Error; err != nil || count != 0 {
		return
	}

	// check if firstAdmin already created or not, continue if firstAdmin not created
	if err = db.Model(&models.User{}).Where("email = ?", admin.Email).First(&models.User{}).Error; err != gorm.ErrRecordNotFound {
		return
	}

	hasherPkg := hasher.New(config.GlobalConfig.Hasher.Cost)
	if admin.Password, err = hasherPkg.HashPassword(cfg.Seed.Admin.Password); err != nil {
		return
	}

	if err = db.Create(&admin).Error; err != nil {
		return
	}

	return
}
