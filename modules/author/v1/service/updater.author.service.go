package service

import (
	"context"
	"gin-starter-gits/config"
	"gin-starter-gits/entity"
	"gin-starter-gits/modules/author/v1/repository"
)

type AuthorUpdater struct {
	cfg        config.Config
	authorRepo repository.AuthorRepositoryUseCase
}

type AuthorUpdaterUseCase interface {
	UpdateAuthor(ctx context.Context, author *entity.Author) error
}

func NewAuthorUpdater(cfg config.Config, authorRepo repository.AuthorRepositoryUseCase) *AuthorUpdater {
	return &AuthorUpdater{
		cfg:        cfg,
		authorRepo: authorRepo,
	}
}

// UpdateAuthor is a function to update author
func (au *AuthorUpdater) UpdateAuthor(ctx context.Context, author *entity.Author) error {
	if err := au.authorRepo.UpdateAuthor(ctx, author); err != nil {
		return err
	}
	return nil
}
