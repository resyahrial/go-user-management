package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/resyahrial/go-user-management/internal/api/middlewares"
)

func (h *Handler) HealthCheck(c *gin.Context) {
	c.Set(middlewares.SuccessKey, map[string]interface{}{
		"message": "OK",
	})
}
