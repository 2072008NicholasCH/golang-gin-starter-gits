package service

import (
	"context"
	"gin-starter-gits/config"
	"gin-starter-gits/entity"
	"gin-starter-gits/modules/author/v1/repository"

	"github.com/google/uuid"
)

type AuthorCreator struct {
	cfg        config.Config
	authorRepo repository.AuthorRepositoryUseCase
}

type AuthorCreatorUseCase interface {
	CreateAuthor(ctx context.Context, name, gender string) (*entity.Author, error)
}

func NewAuthorCreator(cfg config.Config, authorRepo repository.AuthorRepositoryUseCase) *AuthorCreator {
	return &AuthorCreator{
		cfg:        cfg,
		authorRepo: authorRepo,
	}
}

// CreateAuthor is a function to create new author
func (ac *AuthorCreator) CreateAuthor(ctx context.Context, name, gender string) (*entity.Author, error) {
	author := entity.NewAuthor(
		uuid.New(),
		name,
		gender,
		"system",
	)

	if _, err := ac.authorRepo.CreateAuthor(ctx, author); err != nil {
		return nil, err
	}
	return author, nil
}
