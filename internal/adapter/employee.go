package adapter

import (
	"context"

	"github.com/badsector998/go-rbac/internal/domain"
	"github.com/badsector998/go-rbac/internal/lib/transaction"
	"github.com/badsector998/go-rbac/internal/repository"
	"gorm.io/gorm"
)

type EmployeAdapter struct {
	db *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) *EmployeAdapter {
	return &EmployeAdapter{db: db}
}

var _ repository.EmployeeRepository = &EmployeAdapter{}

func (r *EmployeAdapter) Create(ctx context.Context, e domain.Employee) error {
	c := transaction.TransactionFromCtx(ctx, r.db)

	err := c.Save(&e).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *EmployeAdapter) Update(ctx context.Context, e domain.Employee) error {
	c := transaction.TransactionFromCtx(ctx, r.db)

	err := c.Save(&e).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *EmployeAdapter) Delete(ctx context.Context, e domain.Employee) error {
	c := transaction.TransactionFromCtx(ctx, r.db)

	err := c.Delete(&e).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *EmployeAdapter) ReadByID(ctx context.Context, id uint) (domain.Employee, error) {
	c := transaction.TransactionFromCtx(ctx, r.db)

	var e domain.Employee
	err := c.Model(&domain.Employee{}).Find(&e, "id = ?", id).Error
	if err != nil {
		return domain.Employee{}, err
	}

	return e, nil
}

func (r *EmployeAdapter) ReadAll(ctx context.Context) ([]domain.Employee, error) {
	c := transaction.TransactionFromCtx(ctx, r.db)

	var e []domain.Employee
	err := c.Model(&domain.Employee{}).Find(&e).Error
	if err != nil {
		return nil, err
	}

	return e, nil
}
