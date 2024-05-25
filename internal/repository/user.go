package repository

import (
	"context"

	"github.com/badsector998/go-rbac/internal/domain"
)

type User interface {
	Create(ctx context.Context, u domain.User) error
	Update(ctx context.Context, u domain.User) error
	Delete(ctx context.Context, u domain.User) error
	GetUserByEmail(ctx context.Context, email string) (domain.User, error)
	GetUsers(ctx context.Context) ([]domain.User, error)
	GetUserByID(ctx context.Context, id uint) (domain.User, error)
}
