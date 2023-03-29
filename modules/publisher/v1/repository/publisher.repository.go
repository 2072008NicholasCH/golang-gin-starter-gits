package repository

import (
	"context"
	"fmt"
	"gin-starter-gits/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PublisherRepository struct {
	db *gorm.DB
}

type PublisherRepositoryUseCase interface {
	FindAll(ctx context.Context) ([]*entity.Publisher, error)
	FindByID(ctx context.Context, uuid uuid.UUID) (*entity.Publisher, error)
	CreatePublisher(ctx context.Context, publisher *entity.Publisher) (*entity.Publisher, error)
	UpdatePublisher(ctx context.Context, publisher *entity.Publisher) error
	DeletePublisher(ctx context.Context, uuid uuid.UUID) error
}

func NewPublisherRepository(db *gorm.DB) *PublisherRepository {
	return &PublisherRepository{db}
}

func (pr *PublisherRepository) FindAll(ctx context.Context) ([]*entity.Publisher, error) {
	var publishers []*entity.Publisher
	err := pr.db.Find(&publishers).Error
	if err != nil {
		return nil, err
	}
	return publishers, nil
}

func (pr *PublisherRepository) FindByID(ctx context.Context, uuid uuid.UUID) (*entity.Publisher, error) {
	var publisher *entity.Publisher
	err := pr.db.Where("uuid = ?", uuid).First(&publisher).Error
	if err != nil {
		return nil, err
	}
	return publisher, nil
}

// CreatePublisher creates a new publisher
func (pr *PublisherRepository) CreatePublisher(ctx context.Context, publisher *entity.Publisher) (*entity.Publisher, error) {
	err := pr.db.Create(&publisher).Error
	if err != nil {
		return nil, err
	}
	return publisher, nil
}

// UpdatePublisher updates a publisher
//
//	func (pr *PublisherRepository) UpdatePublisher(ctx context.Context, publisher *entity.Publisher) error {
//		err := pr.db.Save(&publisher).Error
//		if err != nil {
//			return err
//		}
//		return nil
//	}
func (pr *PublisherRepository) UpdatePublisher(ctx context.Context, publisher *entity.Publisher) error {
	if err := pr.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		sourceModelPublisher := new(entity.Publisher)
		if err := tx.Clauses(clause.Locking{
			Strength: "UPDATE",
		}).Where("uuid = ?", publisher.UUID).Find(&sourceModelPublisher).Error; err != nil {
			return fmt.Errorf("[PublisherRepository-Update] error while finding publisher")
		}

		if err := tx.Model(&entity.Publisher{}).
			Where("uuid = ?", publisher.UUID).
			UpdateColumns(sourceModelPublisher.MapUpdateFrom(publisher)).
			Error; err != nil {
			return fmt.Errorf("[PublisherRepository-Update] error while updating publisher")
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}

// DeletePublisher deletes a publisher
func (pr *PublisherRepository) DeletePublisher(ctx context.Context, uuid uuid.UUID) error {
	err := pr.db.Where("uuid = ?", uuid).Delete(&entity.Publisher{}).Error
	if err != nil {
		return err
	}
	return nil
}
