package adapter

import (
	"context"

	"github.com/badsector998/go-rbac/internal/domain"
	"github.com/badsector998/go-rbac/internal/lib/transaction"
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
	c := transaction.TransactionFromCtx(ctx, r.db)

	err := c.Save(&b).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *BurjoAdapter) Update(ctx context.Context, b domain.Burjo) error {
	c := transaction.TransactionFromCtx(ctx, r.db)

	err := c.Save(&b).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *BurjoAdapter) Delete(ctx context.Context, b domain.Burjo) error {
	c := transaction.TransactionFromCtx(ctx, r.db)

	err := c.Delete(&b).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *BurjoAdapter) ReadByID(ctx context.Context, id uint) (domain.Burjo, error) {
	c := transaction.TransactionFromCtx(ctx, r.db)

	var b domain.Burjo
	err := c.Model(&domain.Burjo{}).Find(&b, "id = ?", id).Error
	if err != nil {
		return domain.Burjo{}, err
	}

	return b, nil
}

func (r *BurjoAdapter) ReadAll(ctx context.Context) ([]domain.Burjo, error) {
	c := transaction.TransactionFromCtx(ctx, r.db)

	var b []domain.Burjo
	err := c.Model(&domain.Burjo{}).Find(&b).Error
	if err != nil {
		return nil, err
	}

	return b, nil
}
