package repository

import (
	"context"
	"gin-starter-gits/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BookRepository struct {
	db *gorm.DB
}

type BookRepositoryUseCase interface {
	FindAll(ctx context.Context) ([]*entity.Book, error)
	FindByID(ctx context.Context, uuid uuid.UUID) (*entity.Book, error)
}

func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{db}
}

func (r *BookRepository) FindAll(ctx context.Context) ([]*entity.Book, error) {
	var books []*entity.Book
	err := r.db.Find(&books).Error
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (r *BookRepository) FindByID(ctx context.Context, uuid uuid.UUID) (*entity.Book, error) {
	var book *entity.Book
	err := r.db.Where("uuid = ?", uuid).First(&book).Error
	if err != nil {
		return nil, err
	}
	return book, nil
}
