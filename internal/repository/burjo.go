package repository

import (
	"context"

	"github.com/badsector998/go-rbac/domain"
)

type BurjoRepository interface {
	Create(ctx context.Context, b domain.Burjo) error
	Update(ctx context.Context, b domain.Burjo) error
	Delete(ctx context.Context, id uint) error
	ReadByID(ctx context.Context, id uint) (domain.Burjo, error)
	ReadAll(ctx context.Context) ([]domain.Burjo, error)
}
