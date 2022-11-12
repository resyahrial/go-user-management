package testhelper

import (
	"log"
	"os"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DisableLog struct {
}

func (d *DisableLog) Apply(cfg *gorm.Config) error {
	return nil
}

func (d *DisableLog) AfterInitialize(db *gorm.DB) error {
	db.Logger = logger.New(log.New(os.Stdout, "", 0), logger.Config{
		IgnoreRecordNotFoundError: true,
	})
	return nil
}
