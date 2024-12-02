package service

import (
	"OlxScraper/internal/auth"
	"OlxScraper/internal/repository"
)

type Service struct {
	repo  *repository.Repository
	User  UserService
	Admin AdminService
	Auth  AuthService
}

func New(repo *repository.Repository, jwtService auth.JWTService) *Service {
	return &Service{repo: repo, User: NewUserService(repo), Admin: NewAdminService(repo), Auth: NewAuthService(repo, jwtService)}
}
