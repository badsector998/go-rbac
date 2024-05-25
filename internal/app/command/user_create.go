package command

import (
	"context"

	"github.com/badsector998/go-rbac/internal/domain"
	"github.com/badsector998/go-rbac/internal/repository"
)

type UserCreateHandler struct {
	repo repository.User
}

func NewUserCreate(repo repository.User) *UserCreateHandler {
	return &UserCreateHandler{repo: repo}
}

func (h *UserCreateHandler) Handle(ctx context.Context, u domain.User) error {
	return h.repo.Create(ctx, u)
}
