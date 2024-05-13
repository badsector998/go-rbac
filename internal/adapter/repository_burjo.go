package adapter

import (
	"context"

	"github.com/badsector998/go-rbac/domain"
	"github.com/badsector998/go-rbac/internal/repository"
	"gorm.io/gorm"
)

type BurjoAdapter struct {
	db *gorm.DB
}

func NewBurjoRepository(db *gorm.DB) *BurjoAdapter {
	return &BurjoAdapter{db: db}
}

var _ repository.BurjoRepository = &BurjoAdapter{}

func (r *BurjoAdapter) Create(ctx context.Context, b domain.Burjo) error {
	return nil
}

func (r *BurjoAdapter) Update(ctx context.Context, b domain.Burjo) error {
	return nil
}

func (r *BurjoAdapter) Delete(ctx context.Context, id uint) error {
	return nil
}

func (r *BurjoAdapter) ReadByID(ctx context.Context, id uint) (domain.Burjo, error) {
	return domain.Burjo{}, nil
}

func (r *BurjoAdapter) ReadAll(ctx context.Context) ([]domain.Burjo, error) {
	return nil, nil
}
