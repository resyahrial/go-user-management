package seeds

import (
	"github.com/resyahrial/go-user-management/config"
	"gorm.io/gorm"
)

type seedFn func(db *gorm.DB, cfg config.Config) error

var seedList []seedFn = []seedFn{
	seedAdmin,
}

func AutoSeeding(db *gorm.DB, cfg config.Config) {
	var err error
	for _, seed := range seedList {
		if err = seed(db, cfg); err != nil {
			panic(err)
		}
	}
}
