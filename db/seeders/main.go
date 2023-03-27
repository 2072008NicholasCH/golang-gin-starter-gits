package main

import (
	"context"
	"gin-starter-gits/config"
	"gin-starter-gits/entity"
	"gin-starter-gits/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func main() {

	cfg, err := config.LoadConfig(".env")
	checkError(err)

	db, err := utils.NewPostgresGormDB(&cfg.Postgres)
	checkError(err)

	if cfg.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	createrSampleUser(db)
}

func createrSampleUser(db *gorm.DB) {
	userUUID := uuid.New()
	if err := db.WithContext(context.Background()).Model(&entity.User{}).Create(entity.NewUser(
		userUUID,
		"Nicholas CH",
		"nicholas.hanafi@digits.id",
		"test123",
		"system",
	)).Error; err != nil {
		panic(err)
	}
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
