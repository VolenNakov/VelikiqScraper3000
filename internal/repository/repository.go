package repository

import (
	"OlxScraper/internal/db"
)

type Repository struct {
	queries *db.Queries
	User    UserRepository
}

func New(queries *db.Queries) *Repository {
	return &Repository{queries: queries, User: NewUserRepository(queries)}
}
