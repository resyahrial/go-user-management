package pg

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/resyahrial/go-user-management/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var DbInstance *gorm.DB

func InitDatabase(cfg config.Config) *gorm.DB {
	logMode := gormLogger.Silent
	if cfg.App.DebugMode {
		logMode = gormLogger.Info
	}

	var err error
	DbInstance, err = gorm.Open(postgres.Open(cfg.Database.GetDatabaseConnectionString()), &gorm.Config{
		Logger: gormLogger.Default.LogMode(logMode),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: false,
		},
	})
	if err != nil {
		log.Fatalf("failed to establish database connection. error: %s\n", err)
		os.Exit(1)
	}

	sqlDB, err := DbInstance.DB()
	if err != nil {
		fmt.Println(err)
		log.Fatalf("failed to initiate database object. error: %s\n", err)
		os.Exit(1)
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(cfg.Database.MaxIddleConn)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(cfg.Database.MaxOpenConn)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Duration(cfg.Database.ConnMaxLifetime) * time.Hour)

	log.Println("failed to initiate database object")
	return DbInstance
}
