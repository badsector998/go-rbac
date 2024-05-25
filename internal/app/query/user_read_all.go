package query

import (
	"context"
	"fmt"

	"github.com/badsector998/go-rbac/internal/domain"
	"github.com/badsector998/go-rbac/internal/repository"
)

type UserReadAllHandler struct {
	repo repository.User
}

func NewUserReadAll(repo repository.User) *UserReadAllHandler {
	return &UserReadAllHandler{repo: repo}
}

func (h *UserReadAllHandler) Handle(ctx context.Context) ([]domain.User, error) {
	u, err := h.repo.GetUsers(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get all users: %w", err)
	}

	return u, nil
}
