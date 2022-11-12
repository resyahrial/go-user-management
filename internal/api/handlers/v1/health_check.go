package v1

import (
	"github.com/gin-gonic/gin"
	response "github.com/resyahrial/go-user-management/internal/api/handlers/responses"
)

func (h *Handler) HealthCheck(c *gin.Context) {
	res := response.HandleSuccess(map[string]interface{}{
		"message": "OK",
	})
	c.JSON(res.StatusCode, res)
}
