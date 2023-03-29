package service

import (
	"context"
	"gin-starter-gits/config"
	"gin-starter-gits/entity"
	"gin-starter-gits/modules/author/v1/repository"

	"github.com/google/uuid"
)

type AuthorFinder struct {
	afg        config.Config
	authorRepo repository.AuthorRepositoryUseCase
}

type AuthorFinderUseCase interface {
	GetAuthors(ctx context.Context) ([]*entity.Author, error)
	GetAuthorByID(ctx context.Context, uuid uuid.UUID) (*entity.Author, error)
}

func NewAuthorFinder(
	afg config.Config,
	authorRepo repository.AuthorRepositoryUseCase,
) *AuthorFinder {
	return &AuthorFinder{afg, authorRepo}
}

func (s *AuthorFinder) GetAuthors(ctx context.Context) ([]*entity.Author, error) {
	authors, err := s.authorRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	return authors, nil
}

func (s *AuthorFinder) GetAuthorByID(ctx context.Context, uuid uuid.UUID) (*entity.Author, error) {
	author, err := s.authorRepo.FindByID(ctx, uuid)
	if err != nil {
		return nil, err
	}
	return author, nil
}
