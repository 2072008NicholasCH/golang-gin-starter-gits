package builder

import (
	"gin-starter-gits/app"
	"gin-starter-gits/config"
	"gin-starter-gits/modules/publisher/v1/service"

	publisherRepo "gin-starter-gits/modules/publisher/v1/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func BuildPublisherHandler(cfg config.Config, router *gin.Engine, db *gorm.DB) {
	pr := publisherRepo.NewPublisherRepository(db)

	pf := service.NewPublisherFinder(cfg, pr)
	pc := service.NewPublisherCreator(cfg, pr)
	pu := service.NewPublisherUpdater(cfg, pr)
	pd := service.NewPublisherDeleter(cfg, pr)

	app.PublisherFinderHTTPHandler(cfg, router, pf)
	app.PublisherCreatorHTTPHandler(cfg, router, pc)
	app.PublisherUpdaterHTTPHandler(cfg, router, pu, pf)
	app.PublisherDeleterHTTPHandler(cfg, router, pd)
}
