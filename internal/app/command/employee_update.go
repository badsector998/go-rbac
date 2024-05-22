package command

import (
	"context"

	"github.com/badsector998/go-rbac/internal/domain"
	"github.com/badsector998/go-rbac/internal/repository"
)

type EmployeeUpdateHanlder struct {
	repo repository.EmployeeRepository
}

func NewEmployeeUpdate(repo repository.EmployeeRepository) *EmployeeUpdateHanlder {
	return &EmployeeUpdateHanlder{repo: repo}
}

func (h *EmployeeUpdateHanlder) Handle(ctx context.Context, e domain.Employee) error {
	return h.repo.Update(ctx, e)
}
