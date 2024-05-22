package command

import (
	"context"

	"github.com/badsector998/go-rbac/internal/domain"
	"github.com/badsector998/go-rbac/internal/repository"
)

type BurjoUpdateHandler struct {
	repo repository.BurjoRepository
}

func NewBurjoUpdate(repo repository.BurjoRepository) *BurjoUpdateHandler {
	return &BurjoUpdateHandler{repo: repo}
}

func (h *BurjoUpdateHandler) Handle(ctx context.Context, b domain.Burjo) error {
	return h.repo.Update(ctx, b)
}
