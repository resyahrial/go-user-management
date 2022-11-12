package v1

import (
	"github.com/gin-gonic/gin"
)

func CreateUserRoute(router *gin.RouterGroup) {
	router.POST("", func(c *gin.Context) { Handler.CreateUser(c) })
}
