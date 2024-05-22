package repository

import (
	"context"

	"github.com/badsector998/go-rbac/internal/domain"
)

type EmployeeRepository interface {
	Create(ctx context.Context, e domain.Employee) error
	Update(ctx context.Context, e domain.Employee) error
	Delete(ctx context.Context, e domain.Employee) error
	ReadByID(ctx context.Context, id uint) (domain.Employee, error)
	ReadAll(ctx context.Context) ([]domain.Employee, error)
}
