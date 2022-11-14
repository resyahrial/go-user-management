package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/resyahrial/go-user-management/internal/entities"
	"github.com/resyahrial/go-user-management/pkg/exception"
)

var (
	ErrNotLoginYet = exception.NewAuthorizationException().SetMessage("user not login yet")
)

func (m *Middleware) Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		authenticate := &entities.Authentication{}

		if id, ok := c.Get(UserIDKey); !ok {
			handleError(c, ErrNotLoginYet)
			return
		} else if parsedId, ok := id.(string); !ok {
			handleError(c, ErrNotLoginYet)
			return
		} else {
			authenticate.CurrentUserID = parsedId
		}

		if id := c.Param("id"); strings.TrimSpace(id) != "" {
			authenticate.ResourceUserID = id
		}

		if c.Request.Method == http.MethodGet {
			authenticate.Action = "READ"
		} else {
			authenticate.Action = "WRITE"
		}

		// todo: find better way to get module accessed
		if splittedPath := strings.Split(c.Request.URL.Path, "/"); len(splittedPath) >= 4 {
			authenticate.Resource = splittedPath[3]
		}

		err := m.authUsecase.ValidateUserAccess(c.Request.Context(), authenticate)
		if err != nil {
			handleError(c, err)
			return
		}

		c.Next()
	}
}
