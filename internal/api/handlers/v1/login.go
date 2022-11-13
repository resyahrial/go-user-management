package v1

import (
	"github.com/gin-gonic/gin"
	request "github.com/resyahrial/go-user-management/internal/api/handlers/requests"
	response "github.com/resyahrial/go-user-management/internal/api/handlers/responses"
	"github.com/resyahrial/go-user-management/internal/api/middlewares"
	"github.com/resyahrial/go-user-management/internal/entities"
)

func (h *Handler) Login(c *gin.Context) {
	var (
		err   error
		req   *request.LoginRequest
		res   *response.LoginResponse
		login *entities.Login
		token *entities.Token
	)

	if err = c.BindJSON(&req); err != nil {
		c.Set(middlewares.FailureKey, err)
		return
	}

	if login, err = req.CastToLoginEntity(); err != nil {
		c.Set(middlewares.FailureKey, err)
		return
	}

	if token, err = h.authUsecase.Login(c.Request.Context(), login); err != nil {
		c.Set(middlewares.FailureKey, err)
		return
	}

	if res, err = response.NewLoginResponse(token); err != nil {
		c.Set(middlewares.FailureKey, err)
		return
	}

	c.Set(middlewares.SuccessKey, res)
}
