package service

import (
	"OlxScraper/internal/model"
	"OlxScraper/internal/repository"
	"context"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(ctx context.Context, req *model.RegisterRequest) (*model.RegisterResponse, error)
}

type userService struct {
	repo *repository.Repository
}

func NewUserService(repo *repository.Repository) UserService {
	return &userService{repo: repo}
}

func (s *userService) Register(ctx context.Context, req *model.RegisterRequest) (*model.RegisterResponse, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &model.User{
		Email:        req.Email,
		PasswordHash: string(hashedPassword),
	}

	if err := s.repo.User.Create(ctx, user); err != nil {
		return nil, err
	}

	return &model.RegisterResponse{
		ID:    user.ID,
		Email: user.Email,
	}, nil
}
