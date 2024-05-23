package query

import (
	"context"
	"fmt"

	"github.com/badsector998/go-rbac/internal/domain"
	"github.com/badsector998/go-rbac/internal/repository"
)

type BurjoGroupReadAllHandler struct {
	repo repository.BurjoGroupRepository
}

func NewBurjoGroupReadAll(repo repository.BurjoGroupRepository) *BurjoGroupReadAllHandler {
	return &BurjoGroupReadAllHandler{repo: repo}
}

func (h *BurjoGroupReadAllHandler) Handle(ctx context.Context) ([]domain.BurjoGroup, error) {
	bg, err := h.repo.ReadAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get all burjo group: %w", err)
	}

	return bg, nil
}
