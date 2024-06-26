package repository

import (
	"context"

	"github.com/badsector998/go-rbac/internal/domain"
)

//go:generate mockgen -source ./burjo_group.go -destination ./mock/burjo_group_mock.go

type BurjoGroupRepository interface {
	Create(ctx context.Context, bg domain.BurjoGroup) error
	Update(ctx context.Context, bg domain.BurjoGroup) error
	Delete(ctx context.Context, bg domain.BurjoGroup) error
	ReadByID(ctx context.Context, id uint) (domain.BurjoGroup, error)
	ReadAll(ctx context.Context) ([]domain.BurjoGroup, error)
}
