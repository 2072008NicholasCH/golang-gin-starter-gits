package service

import (
	"context"
	"gin-starter-gits/config"
	"gin-starter-gits/entity"
	"gin-starter-gits/modules/publisher/v1/repository"

	"github.com/google/uuid"
)

type PublisherFinder struct {
	pfg           config.Config
	publisherRepo repository.PublisherRepositoryUseCase
}

type PublisherFinderUseCase interface {
	GetPublishers(ctx context.Context) ([]*entity.Publisher, error)
	GetPublisherByID(ctx context.Context, uuid uuid.UUID) (*entity.Publisher, error)
}

func NewPublisherFinder(
	pfg config.Config,
	publisherRepo repository.PublisherRepositoryUseCase,
) *PublisherFinder {
	return &PublisherFinder{pfg, publisherRepo}
}

func (s *PublisherFinder) GetPublishers(ctx context.Context) ([]*entity.Publisher, error) {
	publishers, err := s.publisherRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	return publishers, nil
}

func (s *PublisherFinder) GetPublisherByID(ctx context.Context, uuid uuid.UUID) (*entity.Publisher, error) {
	publisher, err := s.publisherRepo.FindByID(ctx, uuid)
	if err != nil {
		return nil, err
	}
	return publisher, nil
}
