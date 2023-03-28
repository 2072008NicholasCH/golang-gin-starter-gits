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

	app.PublisherFinderHTTPHandler(cfg, router, pf)
}
