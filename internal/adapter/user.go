package adapter

import (
	"context"

	"github.com/badsector998/go-rbac/internal/domain"
	"github.com/badsector998/go-rbac/internal/repository"
	"gorm.io/gorm"
)

type UserAdapter struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserAdapter {
	return &UserAdapter{db: db}
}

var _ repository.User = &UserAdapter{}

func (r *UserAdapter) Create(ctx context.Context, u domain.User) error {
	return nil
}

func (r *UserAdapter) Update(ctx context.Context, u domain.User) error {
	return nil
}

func (r *UserAdapter) Delete(ctx context.Context, u domain.User) error {
	return nil
}

func (r *UserAdapter) GetUserByEmail(ctx context.Context, email string) (domain.User, error) {
	return domain.User{}, nil
}

func (r *UserAdapter) GetUsers(ctx context.Context) ([]domain.User, error) {
	return nil, nil
}

func (r *UserAdapter) GetUserByID(ctx context.Context, id uint) (domain.User, error) {
	return domain.User{}, nil
}
