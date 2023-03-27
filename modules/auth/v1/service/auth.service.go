package service

import (
	"context"
	"gin-starter-gits/config"
	"gin-starter-gits/utils"
	"golang.org/x/crypto/bcrypt"
	"gin-starter-gits/common/errors"
	"gin-starter-gits/entity"
	"gin-starter-gits/modules/auth/v1/repository"
)

// AuthService is a service for auth
type AuthService struct {
	cfg      config.Config
	authRepo repository.AuthRepositoryUseCase
}

// AuthUseCase is a usecase for auth
type AuthUseCase interface {
	// AuthValidate is a function that validates the user
	AuthValidate(ctx context.Context, email, password string) (*entity.User, error)
	// GenerateAccessToken is a function that generates an access token
	GenerateAccessToken(ctx context.Context, user *entity.User) (*entity.Token, error)
}

// NewAuthService is a constructor for AuthService
func NewAuthService(
	cfg config.Config,
	authRepo repository.AuthRepositoryUseCase,
) *AuthService {
	return &AuthService{
		cfg:      cfg,
		authRepo: authRepo,
	}
}

// AuthValidate is a function that validates the user
func (as *AuthService) AuthValidate(ctx context.Context, email, password string) (*entity.User, error) {
	user, err := as.authRepo.GetUserByEmail(ctx, email)

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.ErrWrongLoginCredentials.Error()
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return nil, errors.ErrWrongLoginCredentials.Error()
	}

	return user, nil
}

// GenerateAccessToken is a function that generates an access token
func (as *AuthService) GenerateAccessToken(ctx context.Context, user *entity.User) (*entity.Token, error) {
	token, err := utils.JWTEncode(as.cfg, user.UUID, as.cfg.JWTConfig.Issuer)

	if err != nil {
		return nil, errors.ErrInternalServerError.Error()
	}

	return &entity.Token{
		Token: token,
	}, nil
}
