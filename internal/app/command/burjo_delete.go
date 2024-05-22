package command

import (
	"context"

	"github.com/badsector998/go-rbac/internal/domain"
	"github.com/badsector998/go-rbac/internal/repository"
)

type BurjoDeleteHandler struct {
	repo repository.BurjoRepository
}

func NewBurjoDelete(repo repository.BurjoRepository) *BurjoDeleteHandler {
	return &BurjoDeleteHandler{repo: repo}
}

func (h *BurjoDeleteHandler) Handle(ctx context.Context, b domain.Burjo) error {
	return h.repo.Delete(ctx, b)
}
