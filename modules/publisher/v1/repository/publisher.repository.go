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
	CreatePublisher(ctx context.Context, publisher *entity.Publisher) (*entity.Publisher, error)
	UpdatePublisher(ctx context.Context, publisher *entity.Publisher) (*entity.Publisher, error)
	DeletePublisher(ctx context.Context, id int64) error
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

// CreatePublisher creates a new publisher
func (r *PublisherRepository) CreatePublisher(ctx context.Context, publisher *entity.Publisher) (*entity.Publisher, error) {
	err := r.db.Create(&publisher).Error
	if err != nil {
		return nil, err
	}
	return publisher, nil
}

// UpdatePublisher updates a publisher
func (r *PublisherRepository) UpdatePublisher(ctx context.Context, publisher *entity.Publisher) (*entity.Publisher, error) {
	err := r.db.Save(&publisher).Error
	if err != nil {
		return nil, err
	}
	return publisher, nil
}

// DeletePublisher deletes a publisher
func (r *PublisherRepository) DeletePublisher(ctx context.Context, id int64) error {
	err := r.db.Where("id = ?", id).Delete(&entity.Publisher{}).Error
	if err != nil {
		return err
	}
	return nil
}
