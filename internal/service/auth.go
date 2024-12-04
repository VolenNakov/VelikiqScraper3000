package service

import (
	"OlxScraper/internal/auth"
	"OlxScraper/internal/model"
	"OlxScraper/internal/repository"
	"context"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Login(ctx context.Context, req *model.LoginRequest) (*model.LoginResponse, error)
}

type authService struct {
	repo       *repository.Repository
	jwtService auth.JWTService
}

func NewAuthService(repo *repository.Repository, jwtService auth.JWTService) AuthService {
	return &authService{
		repo:       repo,
		jwtService: jwtService,
	}
}
func (s *authService) Login(ctx context.Context, req *model.LoginRequest) (*model.LoginResponse, error) {
	user, err := s.repo.User.FindByUsername(ctx, req.Username)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password))
	if err != nil {
		return nil, repository.ErrInvalidPassword
	}

	if !user.IsVerified.Bool {
		return nil, repository.ErrUnverifiedUser
	}

	token, err := s.jwtService.CreateToken(user.Role.String)
	if err != nil {
		return nil, repository.ErrInternalError
	}

	return &model.LoginResponse{
		ID:       user.ID,
		Username: user.Username,
		Token:    token,
	}, nil
}
