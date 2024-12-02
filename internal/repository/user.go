package repository

import (
	"OlxScraper/internal/db"
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
	Create(ctx context.Context, user *db.User) (*int64, error)
	FindByEmail(ctx context.Context, email string) (*db.User, error)
	Verify(ctx context.Context, id int) (*bool, error)
	GetUnverifiedUsers(ctx context.Context) ([]db.User, error)
}

type userRepository struct {
	queries *db.Queries
}

func NewUserRepository(queries *db.Queries) UserRepository {
	return &userRepository{queries: queries}
}

func (r *userRepository) Create(ctx context.Context, user *db.User) (*int64, error) {
	userId, err := r.queries.CreateUser(ctx, db.CreateUserParams{
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
	})
	if err != nil {
		if strings.Contains(err.Error(), "UNIQUE constraint failed") {
			return nil, ErrDuplicateEmail
		}
		return nil, err
	}
	return &userId, nil

}
func (r *userRepository) FindByEmail(ctx context.Context, email string) (*db.User, error) {
	user, err := r.queries.GetUserByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrUserNotFound
		}
		fmt.Println(err)
		return nil, ErrInternalError
	}
	return &db.User{
		ID:           user.ID,
		Email:        user.Email,
		IsVerified:   user.IsVerified,
		Role:         user.Role,
		PasswordHash: user.PasswordHash,
	}, nil
}
func (r *userRepository) Verify(ctx context.Context, id int) (*bool, error) {
	isVerified, err := r.queries.VerifyUser(ctx, db.VerifyUserParams{
		IsVerified: sql.NullBool{Bool: true, Valid: true},
		ID:         int64(id),
	})

	if err != nil {
		fmt.Println(err)
		return nil, ErrInternalError
	}

	if !isVerified.Valid && isVerified.Bool {
		return nil, ErrUserNotFound
	}

	return &isVerified.Bool, nil

}

func (r *userRepository) GetUnverifiedUsers(ctx context.Context) ([]db.User, error) {
	users, err := r.queries.GetUnverifiedUsers(ctx)
	if err != nil {
		fmt.Println(err)
		return nil, ErrInternalError
	}
	return users, nil
}
