package command

import (
	"context"

	"github.com/badsector998/go-rbac/internal/domain"
	"github.com/badsector998/go-rbac/internal/repository"
)

type BurjoCreateHandler struct {
	repo repository.BurjoRepository
}

func NewBurjoCreate(repo repository.BurjoRepository) *BurjoCreateHandler {
	return &BurjoCreateHandler{repo: repo}
}

func (h *BurjoCreateHandler) Handle(ctx context.Context, b domain.Burjo) error {
	return nil
}
