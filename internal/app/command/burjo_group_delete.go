package command

import (
	"context"

	"github.com/badsector998/go-rbac/internal/domain"
	"github.com/badsector998/go-rbac/internal/repository"
)

type BurjoGroupDeleteHandler struct {
	repo repository.BurjoGroupRepository
}

func NewBurjoGroupDelete(repo repository.BurjoGroupRepository) *BurjoGroupDeleteHandler {
	return &BurjoGroupDeleteHandler{repo: repo}
}

func (h *BurjoGroupDeleteHandler) Handle(ctx context.Context, bg domain.BurjoGroup) error {
	return h.repo.Delete(ctx, bg)
}
