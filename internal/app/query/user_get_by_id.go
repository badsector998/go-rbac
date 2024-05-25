package query

import (
	"context"
	"fmt"

	"github.com/badsector998/go-rbac/internal/domain"
	"github.com/badsector998/go-rbac/internal/repository"
)

type UserGetByIDHandler struct {
	repo repository.User
}

func NewUserGetByID(repo repository.User) *UserGetByIDHandler {
	return &UserGetByIDHandler{repo: repo}
}

func (h *UserGetByIDHandler) Handle(ctx context.Context, id uint) (domain.User, error) {
	u, err := h.repo.GetUserByID(ctx, id)
	if err != nil {
		return domain.User{}, fmt.Errorf("failed to get user by id: %w", err)
	}

	return u, nil
}
