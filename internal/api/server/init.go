package server

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/resyahrial/go-user-management/config"
	"github.com/resyahrial/go-user-management/internal/api/middlewares"
	route "github.com/resyahrial/go-user-management/internal/api/routes"
	"gorm.io/gorm"
)

func InitGinEngine(cfg config.Config, db *gorm.DB) *gin.Engine {
	var (
		ginMode string
	)

	if cfg.App.DebugMode {
		ginMode = gin.DebugMode
	} else {
		ginMode = gin.ReleaseMode
		gin.DisableConsoleColor()
	}

	gin.SetMode(ginMode)
	r := gin.New()

	customMiddleware := middlewares.NewMiddleware(middlewares.MiddlewareOpts{})

	r.Use(customMiddleware.ResponseWrapper())

	r.Use(gin.CustomRecovery((func(c *gin.Context, recovered interface{}) {
		c.Set(middlewares.FailureKey, fmt.Errorf("panic : %v", recovered))
	})))

	r.NoRoute(func(c *gin.Context) {
		c.Set(middlewares.FailureKey, errors.New("route not found"))
	})

	return route.InitRoutes(r, cfg, db)
}
