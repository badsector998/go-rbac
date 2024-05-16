package adapter

import (
	"context"
	"fmt"

	"github.com/badsector998/go-rbac/internal/domain"
	"github.com/badsector998/go-rbac/internal/lib/transaction"
	"github.com/badsector998/go-rbac/internal/repository"
	"gorm.io/gorm"
)

type BurjoGroupAdapter struct {
	db *gorm.DB
}

func NewBurjoGroupRepository(db *gorm.DB) *BurjoGroupAdapter {
	return &BurjoGroupAdapter{db: db}
}

var _ repository.BurjoGroupRepository = &BurjoGroupAdapter{}

func (r *BurjoGroupAdapter) Create(ctx context.Context, bg domain.BurjoGroup) error {
	c := transaction.TransactionFromCtx(ctx, r.db)

	err := c.Save(&bg).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *BurjoGroupAdapter) Update(ctx context.Context, bg domain.BurjoGroup) error {
	c := transaction.TransactionFromCtx(ctx, r.db)

	if bg.ID == 0 {
		return fmt.Errorf("id is required to update")
	}

	err := c.Save(&bg).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *BurjoGroupAdapter) Delete(ctx context.Context, id uint) error {
	return nil
}

func (r *BurjoGroupAdapter) ReadByID(ctx context.Context, id uint) (domain.BurjoGroup, error) {
	return domain.BurjoGroup{}, nil
}

func (r *BurjoGroupAdapter) ReadAll(ctx context.Context) ([]domain.BurjoGroup, error) {
	return nil, nil
}
