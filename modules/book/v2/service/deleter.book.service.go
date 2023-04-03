package service

import (
	"context"
	"gin-starter-gits/config"
	"gin-starter-gits/modules/book/v1/repository"

	"github.com/google/uuid"
)

type BookDeleter struct {
	cfg      config.Config
	bookRepo repository.BookRepositoryUseCase
}

type BookDeleterUseCase interface {
	DeleteBook(ctx context.Context, uuid uuid.UUID) error
}

func NewBookDeleter(cfg config.Config, bookRepo repository.BookRepositoryUseCase) *BookDeleter {
	return &BookDeleter{cfg, bookRepo}
}

// DeleteBook deletes a book
func (bd *BookDeleter) DeleteBook(ctx context.Context, uuid uuid.UUID) error {
	if err := bd.bookRepo.DeleteBook(ctx, uuid); err != nil {
		return err
	}
	return nil
}
