package query

import (
	"context"
	"fmt"

	"github.com/badsector998/go-rbac/internal/domain"
	"github.com/badsector998/go-rbac/internal/repository"
)

type BurjoGroupReadByIDHandler struct {
	repo repository.BurjoGroupRepository
}

func NewBurjoGroupReadByID(repo repository.BurjoGroupRepository) *BurjoGroupReadByIDHandler {
	return &BurjoGroupReadByIDHandler{repo: repo}
}

func (h *BurjoGroupReadByIDHandler) Handle(ctx context.Context, id uint) (domain.BurjoGroup, error) {
	bg, err := h.repo.ReadByID(ctx, id)
	if err != nil {
		return domain.BurjoGroup{}, fmt.Errorf("failed to get burjo group by id: %w", err)
	}

	return bg, nil
}
