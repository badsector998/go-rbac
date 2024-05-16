package repository

import (
	"context"

	"github.com/badsector998/go-rbac/internal/domain"
)

type UserRepository interface {
	GetUser(ctx context.Context, email string) (domain.User, error)
	GetUsers(ctx context.Context) ([]domain.User, error)
	GetUserByID(ctx context.Context, id uint) (domain.User, error)
	CreateUser(ctx context.Context, u domain.User) error
	DeleteUser(ctx context.Context, id uint) error
	UpdateUser(ctx context.Context, u domain.User) error
}
