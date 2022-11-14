package middlewares

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/resyahrial/go-user-management/internal/entities"
	"github.com/resyahrial/go-user-management/pkg/exception"
)

var (
	UserIDKey = "UserIDKey"
)

func (m *Middleware) Authorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeaders := strings.Fields(c.GetHeader("Authorization"))
		if len(authHeaders) != 2 {
			handleError(c, exception.NewAuthorizationException().SetModule(entities.AuthModule).SetMessage("invalid token, use Bearer <token>"))
			return
		}

		user, err := m.authUsecase.ValidateAccessToken(c.Request.Context(), authHeaders[1])
		if err != nil {
			handleError(c, err)
			return
		}
		c.Set(UserIDKey, user.ID)
		c.Next()
	}
}
