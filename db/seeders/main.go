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

	CreateSampleUser(db)
	CreateSampleAuthor(db)
	CreateSamplePublisher(db)
	CreateSampleBook(db)

}

func CreateSampleUser(db *gorm.DB) {
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

// CreateSampleAuthor create sample author
func CreateSampleAuthor(db *gorm.DB) {
	authorUUID := uuid.New()
	if err := db.WithContext(context.Background()).Model(&entity.Author{}).Create(entity.NewAuthor(
		authorUUID,
		"J.K. Rowling",
		"Female",
		"system",
	)).Error; err != nil {
		panic(err)

	}

	authorUUID2 := uuid.New()
	if err := db.WithContext(context.Background()).Model(&entity.Author{}).Create(entity.NewAuthor(
		authorUUID2,
		"J.R.R. Tolkien",
		"Male",
		"system",
	)).Error; err != nil {
		panic(err)
	}

	authorUUID3 := uuid.New()
	if err := db.WithContext(context.Background()).Model(&entity.Author{}).Create(entity.NewAuthor(
		authorUUID3,
		"George R.R. Martin",
		"Male",
		"system",
	)).Error; err != nil {
		panic(err)
	}
}

// CreateSamplePublisher create sample publisher
func CreateSamplePublisher(db *gorm.DB) {
	publisherUUID := uuid.New()
	if err := db.WithContext(context.Background()).Model(&entity.Publisher{}).Create(entity.NewPublisher(
		publisherUUID,
		"Gramedia",
		"Jakarta",
		"system",
	)).Error; err != nil {
		panic(err)
	}

	publisherUUID2 := uuid.New()
	if err := db.WithContext(context.Background()).Model(&entity.Publisher{}).Create(entity.NewPublisher(
		publisherUUID2,
		"Kompas",
		"Jakarta",
		"system",
	)).Error; err != nil {
		panic(err)
	}

	publisherUUID3 := uuid.New()
	if err := db.WithContext(context.Background()).Model(&entity.Publisher{}).Create(entity.NewPublisher(
		publisherUUID3,
		"HarperCollins",
		"New York",
		"system",
	)).Error; err != nil {
		panic(err)
	}

	publisherUUID4 := uuid.New()
	if err := db.WithContext(context.Background()).Model(&entity.Publisher{}).Create(entity.NewPublisher(
		publisherUUID4,
		"Random House",
		"New York",
		"system",
	)).Error; err != nil {
		panic(err)
	}
}

// CreateSampleBook create sample book
func CreateSampleBook(db *gorm.DB) {
	bookUUID := uuid.New()
	if err := db.WithContext(context.Background()).Model(&entity.Book{}).Create(entity.NewBook(
		bookUUID,
		"Harry Potter and the Philosopher's Stone",
		"9786020302717",
		1, 1,
		"system",
	)).Error; err != nil {
		panic(err)
	}

	bookUUID2 := uuid.New()
	if err := db.WithContext(context.Background()).Model(&entity.Book{}).Create(entity.NewBook(
		bookUUID2,
		"The Hobbit",
		"8580001058174",
		2, 2,
		"system",
	)).Error; err != nil {
		panic(err)
	}

	bookUUID3 := uuid.New()
	if err := db.WithContext(context.Background()).Model(&entity.Book{}).Create(entity.NewBook(
		bookUUID3,
		"A Game of Thrones",
		"9780553593716",
		3, 3,
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
