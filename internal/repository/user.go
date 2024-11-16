package repository

import (
	"OlxScraper/internal/db"
	"OlxScraper/internal/model"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
)

var (
	ErrUserNotFound    = errors.New("user not found")
	ErrDuplicateEmail  = errors.New("email already exists")
	ErrInternalError   = errors.New("internal error")
	ErrInvalidPassword = errors.New("invalid password")
	ErrUnverifiedUser  = errors.New("unverified user")
)

type UserRepository interface {
	Create(ctx context.Context, user *model.User) error
	FindByEmail(ctx context.Context, email string) (*model.User, error)
}

type userRepository struct {
	queries *db.Queries
}

func NewUserRepository(queries *db.Queries) UserRepository {
	return &userRepository{queries: queries}
}

func (r *userRepository) Create(ctx context.Context, user *model.User) error {
	err := r.queries.CreateUser(ctx, db.CreateUserParams{
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
	})
	if err != nil {
		if strings.Contains(err.Error(), "UNIQUE constraint failed") {
			return ErrDuplicateEmail
		}
		return err
	}
	return nil

}
func (r *userRepository) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	user, err := r.queries.GetUser(ctx, email)
	if err != nil {
		fmt.Println(err)
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrUserNotFound
		}
		return nil, ErrInternalError
	}
	return &model.User{
		ID:           user.ID,
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
	}, nil
}
