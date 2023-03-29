package service

import (
	"context"
	"gin-starter-gits/config"
	"gin-starter-gits/entity"
	"gin-starter-gits/modules/publisher/v1/repository"
)

type PublisherCreator struct {
	cfg           config.Config
	publisherRepo repository.PublisherRepositoryUseCase
}

type PublisherCreatorUseCase interface {
	CreatePublisher(ctx context.Context, id int64, name, kota string) (*entity.Publisher, error)
}

func NewPublisherCreator(cfg config.Config, publisherRepo repository.PublisherRepositoryUseCase) *PublisherCreator {
	return &PublisherCreator{cfg, publisherRepo}
}

// CreatePublisher creates a new publisher
func (pc *PublisherCreator) CreatePublisher(ctx context.Context, id int64, name, kota string) (*entity.Publisher, error) {
	publisher := entity.NewPublisher(
		id,
		name,
		kota,
	)

	if _, err := pc.publisherRepo.CreatePublisher(ctx, publisher); err != nil {
		return nil, err
	}
	return publisher, nil
}
