package command

import (
	"context"

	"github.com/badsector998/go-rbac/internal/domain"
	"github.com/badsector998/go-rbac/internal/repository"
)

type BurjoGroupCreateHandler struct {
	repo repository.BurjoGroupRepository
}

func NewBurjoGroupCreate(repo repository.BurjoGroupRepository) *BurjoGroupCreateHandler {
	return &BurjoGroupCreateHandler{repo: repo}
}

func (h *BurjoGroupCreateHandler) Handle(ctx context.Context, bg domain.BurjoGroup) error {
	return h.repo.Create(ctx, bg)
}
