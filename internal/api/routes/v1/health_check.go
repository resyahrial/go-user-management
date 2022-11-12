package v1

import (
	"github.com/gin-gonic/gin"
)

func HealthCheckRoute(router *gin.RouterGroup) {
	router.GET("/health-check", func(c *gin.Context) { Handler.HealthCheck(c) })
}
