package route

import (
	"github.com/gin-gonic/gin"
	v1_handlers "github.com/resyahrial/go-user-management/internal/api/handlers/v1"
	v1_routes "github.com/resyahrial/go-user-management/internal/api/routes/v1"
	"gorm.io/gorm"
)

func InitRoutes(r *gin.Engine, db *gorm.DB) *gin.Engine {
	v1_routes.Handler = v1_handlers.NewHandler(db)
	v1Path := r.Group("/api/v1")
	{
		v1_routes.HealthCheckRoute(v1Path)
	}

	userPath := v1Path.Group("/users")
	{
		v1_routes.CreateUserRoute(userPath)
	}

	return r
}
