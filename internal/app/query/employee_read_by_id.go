package query

import (
	"context"
	"fmt"

	"github.com/badsector998/go-rbac/internal/domain"
	"github.com/badsector998/go-rbac/internal/repository"
)

type EmployeeReadByIDHandler struct {
	repo repository.EmployeeRepository
}

func NewEmployeeReadByID(repo repository.EmployeeRepository) *EmployeeReadByIDHandler {
	return &EmployeeReadByIDHandler{repo: repo}
}

func (h *EmployeeReadByIDHandler) Handle(ctx context.Context, id uint) (domain.Employee, error) {
	e, err := h.repo.ReadByID(ctx, id)
	if err != nil {
		return domain.Employee{}, fmt.Errorf("failed to get employee by id: %w", err)
	}

	return e, nil
}
