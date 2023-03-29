package repository

import (
	"context"
	"fmt"
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
	CreateBook(ctx context.Context, book *entity.Book) (*entity.Book, error)
	UpdateBook(ctx context.Context, book *entity.Book) error
	DeleteBook(ctx context.Context, uuid uuid.UUID) error
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

// CreateBook creates a new book
func (br *BookRepository) CreateBook(ctx context.Context, book *entity.Book) (*entity.Book, error) {
	err := br.db.Create(book).Error
	if err != nil {
		return nil, err
	}
	return book, nil
}

// UpdateBook updates a book
func (br *BookRepository) UpdateBook(ctx context.Context, book *entity.Book) error {
	if err := br.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		sourceModelBook := new(entity.Book)
		if err := tx.Where("uuid = ?", book.UUID).Find(sourceModelBook).Error; err != nil {
			return fmt.Errorf("[BookRepository.UpdateBook] error when finding book")
		}

		if err := tx.Model(&entity.Book{}).Where("uuid = ?", book.UUID).UpdateColumns(sourceModelBook.MapUpdateFrom(book)).Error; err != nil {
			return fmt.Errorf("[BookRepository.UpdateBook] error when updating book")
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}

// DeleteBook deletes a book
func (br *BookRepository) DeleteBook(ctx context.Context, uuid uuid.UUID) error {
	err := br.db.Where("uuid = ?", uuid).Delete(&entity.Book{}).Error
	if err != nil {
		return err
	}
	return nil
}
