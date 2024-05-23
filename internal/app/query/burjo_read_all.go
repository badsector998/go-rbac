package query

import (
	"context"
	"fmt"

	"github.com/badsector998/go-rbac/internal/domain"
	"github.com/badsector998/go-rbac/internal/repository"
)

type BurjoReadAllHandler struct {
	repo repository.BurjoRepository
}

func NewBurjoReadAll(repo repository.BurjoRepository) *BurjoReadAllHandler {
	return &BurjoReadAllHandler{repo: repo}
}

func (h *BurjoReadAllHandler) Handle(ctx context.Context) ([]domain.Burjo, error) {
	b, err := h.repo.ReadAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get all burjo: %w", err)
	}

	return b, nil
}
