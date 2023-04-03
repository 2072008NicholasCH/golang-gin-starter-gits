package builder

import (
	"gin-starter-gits/app"
	"gin-starter-gits/config"

	bookRepo "gin-starter-gits/modules/book/v2/repository"
	"gin-starter-gits/modules/book/v2/service"

	"github.com/gomodule/redigo/redis"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gin-starter-gits/utils"
)

func BuildBookHandler(cfg config.Config, router *gin.Engine, db *gorm.DB, redis *redis.Pool) {
	cache := utils.NewClient(redis)

	ar := bookRepo.NewBookRepository(db, cache)

	bf := service.NewBookFinder(cfg, ar)
	// bc := service.NewBookCreator(cfg, ar)
	// bu := service.NewBookUpdater(cfg, ar)
	// bd := service.NewBookDeleter(cfg, ar)

	app.BookFinderHTTPHandlerV2(cfg, router, bf)
	// app.BookCreatorHTTPHandler(cfg, router, bc)
	// app.BookUpdaterHTTPHandler(cfg, router, bu, bf)
	// app.BookDeleterHTTPHandler(cfg, router, bd)
}
