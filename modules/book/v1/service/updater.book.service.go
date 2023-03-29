package service

import (
	"context"
	"gin-starter-gits/config"
	"gin-starter-gits/entity"
	"gin-starter-gits/modules/book/v1/repository"
)

type BookUpdater struct {
	cfg      config.Config
	bookRepo repository.BookRepositoryUseCase
}

type BookUpdaterUseCase interface {
	UpdateBook(ctx context.Context, book *entity.Book) error
}

func NewBookUpdater(cfg config.Config, bookRepo repository.BookRepositoryUseCase) *BookUpdater {
	return &BookUpdater{cfg, bookRepo}
}

// UpdateBook updates a book
func (bu *BookUpdater) UpdateBook(ctx context.Context, book *entity.Book) error {
	if err := bu.bookRepo.UpdateBook(ctx, book); err != nil {
		return err
	}
	return nil
}
