package repository

import (
	"context"
	"gin-starter-gits/entity"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// AuthRepository is a repository for auth
type AuthRepository struct {
	db *gorm.DB
}

// AuthRepositoryUseCase is a repository for auth
type AuthRepositoryUseCase interface {
	// GetUserByEmail finds a user by email
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
}

// NewAuthRepository returns a auth repository
func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{db}
}

// GetUserByEmail finds a user by email
func (ar *AuthRepository) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	result := new(entity.User)

	if err := ar.db.
		WithContext(ctx).
		Where("email = ?", email).
		Find(result).
		Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, errors.Wrap(err, "[UserRepository-GetUserByEmail] email not found")
	}

	return result, nil
}
