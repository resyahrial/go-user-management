package v1

import (
	"github.com/gin-gonic/gin"
)

func CreateUserRoute(router *gin.RouterGroup) {
	router.POST("", func(c *gin.Context) { Handler.CreateUser(c) })
}

func UpdateUserRoute(router *gin.RouterGroup) {
	router.PUT("/:id", func(c *gin.Context) { Handler.UpdateUser(c) })
}

func GetUserDetailRoute(router *gin.RouterGroup) {
	router.GET("/:id", func(c *gin.Context) { Handler.GetDetail(c) })
}
