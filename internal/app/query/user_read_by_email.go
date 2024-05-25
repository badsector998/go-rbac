package query

import (
	"context"
	"fmt"

	"github.com/badsector998/go-rbac/internal/domain"
	"github.com/badsector998/go-rbac/internal/repository"
)

type UserReadByEmailHandler struct {
	repo repository.User
}

func NewUserReadByEmail(repo repository.User) *UserReadByEmailHandler {
	return &UserReadByEmailHandler{repo: repo}
}

func (h *UserReadByEmailHandler) Handle(ctx context.Context, email string) (domain.User, error) {
	u, err := h.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return domain.User{}, fmt.Errorf("failed to get user by email: %w", err)
	}

	return u, nil
}
