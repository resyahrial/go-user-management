package route

import (
	"github.com/gin-gonic/gin"
	"github.com/resyahrial/go-user-management/config"
	v1_handlers "github.com/resyahrial/go-user-management/internal/api/handlers/v1"
	v1_routes "github.com/resyahrial/go-user-management/internal/api/routes/v1"
	"gorm.io/gorm"
)

func InitRoutes(r *gin.Engine, cfg config.Config, db *gorm.DB) *gin.Engine {
	v1_routes.Handler = v1_handlers.NewHandler(cfg, db)
	v1Path := r.Group("/api/v1")
	{
		v1_routes.HealthCheckRoute(v1Path)
		v1_routes.LoginRoute(v1Path)
	}

	userPath := v1Path.Group("/users")
	{
		v1_routes.CreateUserRoute(userPath)
		v1_routes.UpdateUserRoute(userPath)
		v1_routes.GetUserDetailRoute(userPath)
		v1_routes.GetUserListRoute(userPath)
		v1_routes.DeleteUserRoute(userPath)
	}

	return r
}
