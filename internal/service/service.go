package service

import (
	"OlxScraper/internal/repository"
)

type Service struct {
	repo *repository.Repository
	User UserService
}

func New(repo *repository.Repository) *Service {
	return &Service{repo: repo, User: NewUserService(repo)}
}
