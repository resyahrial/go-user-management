package v1

import (
	"github.com/gin-gonic/gin"
	request "github.com/resyahrial/go-user-management/internal/api/handlers/requests"
	response "github.com/resyahrial/go-user-management/internal/api/handlers/responses"
	"github.com/resyahrial/go-user-management/internal/entities"
)

func (h *Handler) CreateUser(c *gin.Context) {
	var (
		err  error
		req  *request.CreateUserRequest
		user *entities.User
	)

	if err = c.BindJSON(&req); err != nil {
		res := response.HandleError(c, err)
		c.JSON(res.StatusCode, res)
		return
	}

	if user, err = req.CastToUserEntity(); err != nil {
		res := response.HandleError(c, err)
		c.JSON(res.StatusCode, res)
		return
	}

	if user, err = h.userUsecase.CreateUser(c.Request.Context(), user); err != nil {
		res := response.HandleError(c, err)
		c.JSON(res.StatusCode, res)
		return
	}

	res := response.HandleSuccess(user)
	c.JSON(res.StatusCode, res)
}
