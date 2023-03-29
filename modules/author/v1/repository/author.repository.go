package repository

import (
	"context"
	"fmt"
	"gin-starter-gits/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type AuthorRepository struct {
	db *gorm.DB
}

type AuthorRepositoryUseCase interface {
	FindAll(ctx context.Context) ([]*entity.Author, error)
	FindByID(ctx context.Context, uuid uuid.UUID) (*entity.Author, error)
	CreateAuthor(ctx context.Context, author *entity.Author) (*entity.Author, error)
	UpdateAuthor(ctx context.Context, author *entity.Author) error
	DeleteAuthor(ctx context.Context, uuid uuid.UUID) error
}

func NewAuthorRepository(db *gorm.DB) *AuthorRepository {
	return &AuthorRepository{db}
}

func (ar *AuthorRepository) FindAll(ctx context.Context) ([]*entity.Author, error) {
	var authors []*entity.Author
	err := ar.db.Find(&authors).Error
	if err != nil {
		return nil, err
	}
	return authors, nil
}

func (ar *AuthorRepository) FindByID(ctx context.Context, uuid uuid.UUID) (*entity.Author, error) {
	var author *entity.Author
	err := ar.db.Where("uuid = ?", uuid).First(&author).Error
	if err != nil {
		return nil, err
	}
	return author, nil
}

// CreateAuthor creates a new author
func (ar *AuthorRepository) CreateAuthor(ctx context.Context, author *entity.Author) (*entity.Author, error) {
	err := ar.db.Create(&author).Error
	if err != nil {
		return nil, err
	}
	return author, nil
}

// UpdateAuthor updates an existing author
func (ar *AuthorRepository) UpdateAuthor(ctx context.Context, author *entity.Author) error {
	if err := ar.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		sourceModelAuthor := new(entity.Author)
		if err := tx.Clauses(clause.Locking{
			Strength: "UPDATE",
		}).Where("uuid = ?", author.UUID).Find(&sourceModelAuthor).Error; err != nil {
			return fmt.Errorf("[AuthorRepository-Update] error while finding author")
		}

		if err := tx.Model(&entity.Author{}).
			Where("uuid = ?", author.UUID).
			UpdateColumns(sourceModelAuthor.MapUpdateFrom(author)).
			Error; err != nil {
			return fmt.Errorf("[AuthorRepository-Update] error while updating author")
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}

// DeleteAuthor deletes an existing author
func (ar *AuthorRepository) DeleteAuthor(ctx context.Context, uuid uuid.UUID) error {
	err := ar.db.Where("uuid = ?", uuid).Delete(&entity.Author{}).Error
	if err != nil {
		return err
	}

	return nil
}
