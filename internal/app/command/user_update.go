package command

import (
	"context"

	"github.com/badsector998/go-rbac/internal/domain"
	"github.com/badsector998/go-rbac/internal/repository"
)

type UserUpdateHandler struct {
	repo repository.User
}

func NewUserUpdate(repo repository.User) *UserUpdateHandler {
	return &UserUpdateHandler{repo: repo}
}

func (h *UserUpdateHandler) Handle(ctx context.Context, u domain.User) error {
	return h.repo.Update(ctx, u)
}
