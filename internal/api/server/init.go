package server

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/resyahrial/go-user-management/config"
	response "github.com/resyahrial/go-user-management/internal/api/handlers/responses"
)

func InitGinEngine(cfg config.Config) *gin.Engine {
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
	r := gin.Default()

	// r.Use(middleware.RequestID())
	// r.Use(middleware.BeforeAfterRequest())
	// otel.SetTracerProvider(tp)
	// r.Use(otelgin.Middleware(appConfig.Name, otelgin.WithTracerProvider(tp)))

	r.Use(gin.CustomRecovery((func(c *gin.Context, recovered interface{}) {
		err := fmt.Errorf("panic : %v", recovered)
		res := response.HandleError(c, err)
		c.AbortWithStatusJSON(res.StatusCode, res)
	})))

	r.NoRoute(func(c *gin.Context) {
		err := errors.New("route not found")
		res := response.HandleError(c, err)
		c.JSON(res.StatusCode, res)
	})

	return r
}
