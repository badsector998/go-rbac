package repository

import (
	"context"

	"github.com/badsector998/go-rbac/domain"
)

type EmployeeRepository interface {
	Create(ctx context.Context, b domain.Employee) error
	Update(ctx context.Context, b domain.Employee) error
	Delete(ctx context.Context, id uint) error
	ReadByID(ctx context.Context, id uint) (domain.Employee, error)
	ReadAll(ctx context.Context) ([]domain.Employee, error)
}
