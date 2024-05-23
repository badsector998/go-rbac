package query

import (
	"context"
	"fmt"

	"github.com/badsector998/go-rbac/internal/domain"
	"github.com/badsector998/go-rbac/internal/repository"
)

type BurjoReadByIDHandler struct {
	repo repository.BurjoRepository
}

func NewBurjoReadByID(repo repository.BurjoRepository) *BurjoReadByIDHandler {
	return &BurjoReadByIDHandler{repo: repo}
}

func (h *BurjoReadByIDHandler) Handle(ctx context.Context, id uint) (domain.Burjo, error) {
	b, err := h.repo.ReadByID(ctx, id)
	if err != nil {
		return domain.Burjo{}, fmt.Errorf("failed to get burjo by id: %w", err)
	}

	return b, nil
}
