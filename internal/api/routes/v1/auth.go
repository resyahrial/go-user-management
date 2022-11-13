package v1

import (
	"github.com/gin-gonic/gin"
)

func LoginRoute(router *gin.RouterGroup) {
	router.POST("", func(c *gin.Context) { Handler.Login(c) })
}
