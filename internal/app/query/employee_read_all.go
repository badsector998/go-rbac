package query

import (
	"context"
	"fmt"

	"github.com/badsector998/go-rbac/internal/domain"
	"github.com/badsector998/go-rbac/internal/repository"
)

type EmployeeReadAllHandler struct {
	repo repository.EmployeeRepository
}

func NewEmployeeReadAll(repo repository.EmployeeRepository) *EmployeeReadAllHandler {
	return &EmployeeReadAllHandler{repo: repo}
}

func (h *EmployeeReadAllHandler) Handle(ctx context.Context) ([]domain.Employee, error) {
	e, err := h.repo.ReadAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get all employee: %w", err)
	}

	return e, nil
}
