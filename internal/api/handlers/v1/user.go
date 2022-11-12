package v1

import (
	"github.com/gin-gonic/gin"
	request "github.com/resyahrial/go-user-management/internal/api/handlers/requests"
	response "github.com/resyahrial/go-user-management/internal/api/handlers/responses"
	"github.com/resyahrial/go-user-management/internal/api/middlewares"
	"github.com/resyahrial/go-user-management/internal/entities"
)

func (h *Handler) CreateUser(c *gin.Context) {
	var (
		err  error
		req  *request.CreateUserRequest
		res  *response.UserResponse
		user *entities.User
	)

	if err = c.BindJSON(&req); err != nil {
		c.Set(middlewares.FailureKey, err)
		return
	}

	if user, err = req.CastToUserEntity(); err != nil {
		c.Set(middlewares.FailureKey, err)
		return
	}

	if user, err = h.userUsecase.Create(c.Request.Context(), user); err != nil {
		c.Set(middlewares.FailureKey, err)
		return
	}

	if res, err = response.NewUserResponse(user); err != nil {
		c.Set(middlewares.FailureKey, err)
		return
	}

	c.Set(middlewares.SuccessKey, res)
}

func (h *Handler) UpdateUser(c *gin.Context) {
	var (
		err  error
		req  *request.UpdateUserRequest
		res  *response.UserResponse
		user *entities.User
	)

	if err = c.BindJSON(&req); err != nil {
		c.Set(middlewares.FailureKey, err)
		return
	}

	if user, err = req.CastToUserEntity(); err != nil {
		c.Set(middlewares.FailureKey, err)
		return
	}

	userId := c.Param("id")
	if user, err = h.userUsecase.Update(c.Request.Context(), userId, user); err != nil {
		c.Set(middlewares.FailureKey, err)
		return
	}

	if res, err = response.NewUserResponse(user); err != nil {
		c.Set(middlewares.FailureKey, err)
		return
	}

	c.Set(middlewares.SuccessKey, res)
}

func (h *Handler) GetDetail(c *gin.Context) {
	var (
		err  error
		res  *response.UserResponse
		user *entities.User
	)

	// userId := c.Param("id")
	// if user, err = h.userUsecase.GetDetail(c.Request.Context(), userId); err != nil {
	// 	c.Set(middlewares.FailureKey, err)
	// 	return
	// }

	if res, err = response.NewUserResponse(user); err != nil {
		c.Set(middlewares.FailureKey, err)
		return
	}

	c.Set(middlewares.SuccessKey, res)
}
