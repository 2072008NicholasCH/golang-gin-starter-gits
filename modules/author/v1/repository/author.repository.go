package repository

import (
	"context"
	"gin-starter-gits/entity"
	"gorm.io/gorm"
)

type AuthorRepository struct {
	db *gorm.DB
}

type AuthorRepositoryInterface interface {
	FindAll(ctx context.Context) ([]*entity.Author, error)
	FindByID(ctx context.Context, id int) (*entity.Author, error)
}

func NewAuthorRepository(db *gorm.DB) *AuthorRepository {
	return &AuthorRepository{db}
}

func (r *AuthorRepository) FindAll(ctx context.Context) ([]*entity.Author, error) {
	var authors []*entity.Author
	err := r.db.WithContext(ctx).Find(&authors).Error
	if err != nil {
		return nil, err
	}
	return authors, nil
}

func (r *AuthorRepository) FindByID(ctx context.Context, id int) (*entity.Author, error) {
	var author *entity.Author
	err := r.db.WithContext(ctx).First(&author, id).Error
	if err != nil {
		return nil, err
	}
	return author, nil
}
