package repository

import (
	"context"
	"gin-starter-gits/entity"

	"gorm.io/gorm"
)

type BookRepository struct {
	db *gorm.DB
}

type BookRepositoryUseCase interface {
	FindAll(ctx context.Context) ([]*entity.Book, error)
	FindByID(ctx context.Context, id int64) (*entity.Book, error)
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

func (r *BookRepository) FindByID(ctx context.Context, id int64) (*entity.Book, error) {
	var book *entity.Book
	err := r.db.Where("id = ?", id).First(&book).Error
	if err != nil {
		return nil, err
	}
	return book, nil
}
