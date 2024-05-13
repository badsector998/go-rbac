package repository

import (
	"context"

	"github.com/badsector998/go-rbac/domain"
)

type BurjoGroupRepository interface {
	Create(ctx context.Context, bg domain.BurjoGroup) error
	Update(ctx context.Context, bg domain.BurjoGroup) error
	Delete(ctx context.Context, id uint) error
	ReadByID(ctx context.Context, id uint) (domain.BurjoGroup, error)
	ReadAll(ctx context.Context) ([]domain.BurjoGroup, error)
}
