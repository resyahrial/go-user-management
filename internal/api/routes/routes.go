package route

import (
	"github.com/gin-gonic/gin"
	"github.com/resyahrial/go-user-management/config"
	v1_handlers "github.com/resyahrial/go-user-management/internal/api/handlers/v1"
	"github.com/resyahrial/go-user-management/internal/api/middlewares"
	v1_routes "github.com/resyahrial/go-user-management/internal/api/routes/v1"
	"gorm.io/gorm"
)

const (
	// path
	base = "/api/v1"
	user = "/users"
)

func InitRoutes(r *gin.Engine, cfg config.Config, db *gorm.DB) *gin.Engine {
	customMiddleware := middlewares.NewMiddleware(middlewares.MiddlewareOpts{
		Db:  db,
		Cfg: cfg,
	}, middlewares.WithAuthUsecase)
	v1_routes.Handler = v1_handlers.NewHandler(cfg, db)

	v1Path := r.Group(base)
	{
		v1_routes.HealthCheckRoute(v1Path)
		v1_routes.LoginRoute(v1Path)
	}

	userPath := v1Path.Group(user)
	{
		v1_routes.CreateUserRoute(userPath)
	}

	userPathWithAuth := v1Path.Group(user, customMiddleware.Authorization())
	{
		v1_routes.UpdateUserRoute(userPathWithAuth)
		v1_routes.GetUserDetailRoute(userPathWithAuth)
		v1_routes.GetUserListRoute(userPathWithAuth)
		v1_routes.DeleteUserRoute(userPathWithAuth)
	}

	return r
}
