package command

import (
	"context"

	"github.com/badsector998/go-rbac/internal/domain"
	"github.com/badsector998/go-rbac/internal/repository"
)

type EmployeeCreateHandler struct {
	repo repository.EmployeeRepository
}

func NewEmployeeCreate(repo repository.EmployeeRepository) *EmployeeCreateHandler {
	return &EmployeeCreateHandler{repo: repo}
}

func (h *EmployeeCreateHandler) Handle(ctx context.Context, e domain.Employee) error {
	return h.repo.Create(ctx, e)
}
