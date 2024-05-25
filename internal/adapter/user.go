package adapter

import (
	"context"

	"github.com/badsector998/go-rbac/internal/domain"
	"github.com/badsector998/go-rbac/internal/lib/transaction"
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
	c := transaction.TransactionFromCtx(ctx, r.db)

	err := c.Save(&u).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *UserAdapter) Update(ctx context.Context, u domain.User) error {
	c := transaction.TransactionFromCtx(ctx, r.db)

	err := c.Save(&u).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *UserAdapter) Delete(ctx context.Context, u domain.User) error {
	c := transaction.TransactionFromCtx(ctx, r.db)

	err := c.Delete(&u).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *UserAdapter) GetUserByEmail(ctx context.Context, email string) (domain.User, error) {
	c := transaction.TransactionFromCtx(ctx, r.db)

	var u domain.User
	err := c.Where("email = ?", email).First(&domain.User{}).Error
	if err != nil {
		return domain.User{}, err
	}

	return u, nil
}

func (r *UserAdapter) GetUsers(ctx context.Context) ([]domain.User, error) {
	c := transaction.TransactionFromCtx(ctx, r.db)

	var users []domain.User
	err := c.Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserAdapter) GetUserByID(ctx context.Context, id uint) (domain.User, error) {
	c := transaction.TransactionFromCtx(ctx, r.db)

	var u domain.User
	err := c.Where("id = ?", id).First(&u).Error
	if err != nil {
		return domain.User{}, err
	}

	return u, nil
}
