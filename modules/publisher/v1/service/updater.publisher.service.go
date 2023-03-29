package service

import (
	"context"
	"gin-starter-gits/common/errors"
	"gin-starter-gits/config"
	"gin-starter-gits/entity"
	"gin-starter-gits/modules/publisher/v1/repository"
)

type PublisherUpdater struct {
	cfg           config.Config
	publisherRepo repository.PublisherRepositoryUseCase
}

type PublisherUpdaterUseCase interface {
	UpdatePublisher(ctx context.Context, publisher *entity.Publisher) error
}

func NewPublisherUpdater(cfg config.Config, publisherRepo repository.PublisherRepositoryUseCase) *PublisherUpdater {
	return &PublisherUpdater{cfg, publisherRepo}
}

func (pu *PublisherUpdater) UpdatePublisher(ctx context.Context, publisher *entity.Publisher) error {
	if err := pu.publisherRepo.UpdatePublisher(ctx, publisher); err != nil {
		return errors.ErrInternalServerError.Error()
	}
	return nil
}
