package command

import (
	"context"

	"github.com/badsector998/go-rbac/internal/domain"
	"github.com/badsector998/go-rbac/internal/repository"
)

type EmployeeDeleteHandler struct {
	repo repository.EmployeeRepository
}

func NewEmployeeDelete(repo repository.EmployeeRepository) *EmployeeDeleteHandler {
	return &EmployeeDeleteHandler{repo: repo}
}

func (h *EmployeeDeleteHandler) Handle(ctx context.Context, e domain.Employee) error {
	return h.repo.Delete(ctx, e)
}
