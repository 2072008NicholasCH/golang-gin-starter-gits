package builder

import (
	"gin-starter-gits/app"
	"gin-starter-gits/config"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	authRepo "gin-starter-gits/modules/auth/v1/repository"
	auth "gin-starter-gits/modules/auth/v1/service"
)

// BuildAuthHandler build auth handlers
// starting from handler down to repository or tool.
func BuildAuthHandler(
	cfg config.Config,
	router *gin.Engine,
	db *gorm.DB,
	// redisPool *redis.Pool,
	// awsSession *session.Session
) {
	// Repository
	ar := authRepo.NewAuthRepository(db)

	uc := auth.NewAuthService(cfg, ar)

	app.AuthHTTPHandler(cfg, router, uc)
}
