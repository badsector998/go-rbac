package command

import (
	"context"

	"github.com/badsector998/go-rbac/internal/domain"
	"github.com/badsector998/go-rbac/internal/repository"
)

type UserDeleteHandler struct {
	repo repository.User
}

func NewUserDelete(repo repository.User) *UserDeleteHandler {
	return &UserDeleteHandler{repo: repo}
}

func (h *UserDeleteHandler) Handle(ctx context.Context, u domain.User) error {
	return h.repo.Delete(ctx, u)
}
