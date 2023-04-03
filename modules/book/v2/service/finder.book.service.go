package service

import (
	"context"
	"gin-starter-gits/config"
	"gin-starter-gits/entity"
	repository "gin-starter-gits/modules/book/v2/repository"

	"github.com/google/uuid"
)

type BookFinder struct {
	bfg      config.Config
	bookRepo repository.BookRepositoryUseCase
}

type BookFinderUseCase interface {
	GetBooks(ctx context.Context) ([]*entity.Book, error)
	GetBookByID(ctx context.Context, uuid uuid.UUID) (*entity.Book, error)
}

func NewBookFinder(
	bfg config.Config,
	bookRepo repository.BookRepositoryUseCase,
) *BookFinder {
	return &BookFinder{bfg, bookRepo}
}

func (s *BookFinder) GetBooks(ctx context.Context) ([]*entity.Book, error) {
	books, err := s.bookRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (s *BookFinder) GetBookByID(ctx context.Context, uuid uuid.UUID) (*entity.Book, error) {
	book, err := s.bookRepo.FindByID(ctx, uuid)
	if err != nil {
		return nil, err
	}
	return book, nil
}
