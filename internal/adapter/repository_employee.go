package adapter

import (
	"context"

	"github.com/badsector998/go-rbac/domain"
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
	return nil
}

func (r *EmployeAdapter) Update(ctx context.Context, e domain.Employee) error {
	return nil
}

func (r *EmployeAdapter) Delete(ctx context.Context, id uint) error {
	return nil
}

func (r *EmployeAdapter) ReadByID(ctx context.Context, id uint) (domain.Employee, error) {
	return domain.Employee{}, nil
}

func (r *EmployeAdapter) ReadAll(ctx context.Context) ([]domain.Employee, error) {
	return nil, nil
}
