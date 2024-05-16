package adapter

import (
	"context"
	"testing"

	"github.com/badsector998/go-rbac/internal/app/command"
	"github.com/badsector998/go-rbac/internal/domain"
	mock_repository "github.com/badsector998/go-rbac/internal/repository/mock"
	"go.uber.org/mock/gomock"
)

func TestBurjoCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	burjoCreate := mock_repository.NewMockBurjoGroupRepository(ctrl)
	a := command.NewBurjoGroupCreate(burjoCreate)

	bg := domain.BurjoGroup{
		Name:        "Group Jogja",
		Owner:       "Dadang Rokes",
		Description: "Burjo Group di Jogja",
	}

	burjoCreate.EXPECT().Create(context.Background(), bg).Return(nil)

	err := a.Handle(context.Background(), bg)
	if err != nil {
		t.Errorf("failed")
	}

}
