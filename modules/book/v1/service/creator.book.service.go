package service

import (
	"context"
	"gin-starter-gits/config"
	"gin-starter-gits/entity"
	"gin-starter-gits/modules/book/v1/repository"

	"github.com/google/uuid"
)

type BookCreator struct {
	cfg      config.Config
	bookRepo repository.BookRepositoryUseCase
}

type BookCreatorUseCase interface {
	CreateBook(ctx context.Context, isbn string, title string, authorID int64, publisherID int64) (*entity.Book, error)
}

func NewBookCreator(cfg config.Config, bookRepo repository.BookRepositoryUseCase) *BookCreator {
	return &BookCreator{cfg, bookRepo}
}

// CreateBook creates a new book
func (bc *BookCreator) CreateBook(ctx context.Context, isbn string, title string, authorID int64, publisherID int64) (*entity.Book, error) {
	book := entity.NewBook(
		uuid.New(),
		isbn,
		title,
		authorID,
		publisherID,
		"system",
	)

	if _, err := bc.bookRepo.CreateBook(ctx, book); err != nil {
		return nil, err
	}
	return book, nil
}
