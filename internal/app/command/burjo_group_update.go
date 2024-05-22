package command

import (
	"context"

	"github.com/badsector998/go-rbac/internal/domain"
	"github.com/badsector998/go-rbac/internal/repository"
)

type BurjoGroupUpdateHandler struct {
	repo repository.BurjoGroupRepository
}

func NewBurjoGroupUpdate(repo repository.BurjoGroupRepository) *BurjoGroupUpdateHandler {
	return &BurjoGroupUpdateHandler{repo: repo}
}

func (h *BurjoGroupUpdateHandler) Handle(ctx context.Context, bg domain.BurjoGroup) error {
	return h.repo.Update(ctx, bg)
}
