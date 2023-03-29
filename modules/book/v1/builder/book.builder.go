package builder

import (
	"gin-starter-gits/app"
	"gin-starter-gits/config"

	bookRepo "gin-starter-gits/modules/book/v1/repository"
	"gin-starter-gits/modules/book/v1/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func BuildBookHandler(cfg config.Config, router *gin.Engine, db *gorm.DB) {
	ar := bookRepo.NewBookRepository(db)

	bf := service.NewBookFinder(cfg, ar)
	bc := service.NewBookCreator(cfg, ar)
	bu := service.NewBookUpdater(cfg, ar)
	bd := service.NewBookDeleter(cfg, ar)

	app.BookFinderHTTPHandler(cfg, router, bf)
	app.BookCreatorHTTPHandler(cfg, router, bc)
	app.BookUpdaterHTTPHandler(cfg, router, bu, bf)
	app.BookDeleterHTTPHandler(cfg, router, bd)
}
