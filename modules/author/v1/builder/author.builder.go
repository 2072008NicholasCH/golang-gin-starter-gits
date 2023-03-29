package builder

import (
	"gin-starter-gits/app"
	"gin-starter-gits/config"

	authorRepo "gin-starter-gits/modules/author/v1/repository"
	"gin-starter-gits/modules/author/v1/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func BuildAuthorHandler(cfg config.Config, router *gin.Engine, db *gorm.DB) {
	ar := authorRepo.NewAuthorRepository(db)

	af := service.NewAuthorFinder(cfg, ar)
	ac := service.NewAuthorCreator(cfg, ar)
	au := service.NewAuthorUpdater(cfg, ar)
	ad := service.NewAuthorDeleter(cfg, ar)

	app.AuthorFinderHTTPHandler(cfg, router, af)
	app.AuthorCreatorHTTPHandler(cfg, router, ac)
	app.AuthorUpdaterHTTPHandler(cfg, router, au, af)
	app.AuthorDeleterHTTPHandler(cfg, router, ad)

}
