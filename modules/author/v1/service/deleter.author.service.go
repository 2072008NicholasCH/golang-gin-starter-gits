package service

import (
	"context"
	"gin-starter-gits/config"
	"gin-starter-gits/modules/author/v1/repository"

	"github.com/google/uuid"
)

type AuthorDeleter struct {
	cfg        config.Config
	authorRepo repository.AuthorRepositoryUseCase
}

type AuthorDeleterUseCase interface {
	DeleteAuthor(ctx context.Context, uuid uuid.UUID) error
}

func NewAuthorDeleter(cfg config.Config, authorRepo repository.AuthorRepositoryUseCase) *AuthorDeleter {
	return &AuthorDeleter{
		cfg:        cfg,
		authorRepo: authorRepo,
	}
}

// DeleteAuthor is a function to delete author
func (ad *AuthorDeleter) DeleteAuthor(ctx context.Context, uuid uuid.UUID) error {
	if err := ad.authorRepo.DeleteAuthor(ctx, uuid); err != nil {
		return err
	}
	return nil
}
