package repository

import (
	"context"
	"gin-starter-gits/entity"

	"gorm.io/gorm"
)

type PublisherRepository struct {
	db *gorm.DB
}

type PublisherRepositoryUseCase interface {
	FindAll(ctx context.Context) ([]*entity.Publisher, error)
	FindByID(ctx context.Context, id int64) (*entity.Publisher, error)
}

func NewPublisherRepository(db *gorm.DB) *PublisherRepository {
	return &PublisherRepository{db}
}

func (r *PublisherRepository) FindAll(ctx context.Context) ([]*entity.Publisher, error) {
	var publishers []*entity.Publisher
	err := r.db.Find(&publishers).Error
	if err != nil {
		return nil, err
	}
	return publishers, nil
}

func (r *PublisherRepository) FindByID(ctx context.Context, id int64) (*entity.Publisher, error) {
	var publisher *entity.Publisher
	err := r.db.Where("id = ?", id).First(&publisher).Error
	if err != nil {
		return nil, err
	}
	return publisher, nil
}
