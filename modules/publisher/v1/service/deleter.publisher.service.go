package service

import (
	"context"
	"gin-starter-gits/config"
	"gin-starter-gits/modules/publisher/v1/repository"

	"github.com/google/uuid"
)

type PublisherDeleter struct {
	cfg           config.Config
	publisherRepo repository.PublisherRepositoryUseCase
}

type PublisherDeleterUseCase interface {
	DeletePublisher(ctx context.Context, uuid uuid.UUID) error
}

func NewPublisherDeleter(cfg config.Config, publisherRepo repository.PublisherRepositoryUseCase) *PublisherDeleter {
	return &PublisherDeleter{cfg, publisherRepo}
}

func (pd *PublisherDeleter) DeletePublisher(ctx context.Context, uuid uuid.UUID) error {
	if err := pd.publisherRepo.DeletePublisher(ctx, uuid); err != nil {
		return err
	}
	return nil
}
