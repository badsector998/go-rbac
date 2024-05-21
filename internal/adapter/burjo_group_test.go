package adapter

import (
	"context"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/badsector998/go-rbac/internal/domain"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestBurjoGroup(t *testing.T) {
	s := new(TestSuite)
	s.Setup(t)

	tt := []struct {
		name     string
		expectFn func()
		actualFn func(context.Context) error
		wantErr  error
	}{
		{
			name: "read",
			expectFn: func() {
				s.mockClient.ExpectQuery(`SELECT * FROM "burjo_groups" WHERE id = $1 AND "burjo_groups"."deleted_at" IS NULL`).
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "name", "owner", "description"}).AddRow(
						1, time.Now(), time.Now(), gorm.DeletedAt{}, "a", "b", "c",
					))
			},
			actualFn: func(ctx context.Context) error {
				bgr := NewBurjoGroupRepository(s.client)

				_, err := bgr.ReadByID(ctx, 1)
				if err != nil {
					return err
				}

				return nil
			},
			wantErr: nil,
		},
		{
			name: "create",
			expectFn: func() {
				s.mockClient.ExpectQuery(`INSERT INTO "burjo_groups" ("created_at","updated_at","deleted_at","name","owner","description") VALUES ($1,$2,$3,$4,$5,$6) RETURNING "id"`).
					WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), "a", "b", "c").
					WillReturnRows(sqlmock.NewRows([]string{"id"}))
			},
			actualFn: func(ctx context.Context) error {
				bgr := NewBurjoGroupRepository(s.client)

				bg := domain.BurjoGroup{
					Name:        "a",
					Owner:       "b",
					Description: "c",
				}

				err := bgr.Create(ctx, bg)
				if err != nil {
					return err
				}

				return nil
			},
			wantErr: nil,
		},
		{
			name: "update",
			expectFn: func() {
				s.mockClient.ExpectExec(`UPDATE "burjo_groups" SET "created_at"=$1,"updated_at"=$2,"deleted_at"=$3,"name"=$4,"owner"=$5,"description"=$6 WHERE "burjo_groups"."deleted_at" IS NULL AND "id" = $7`).
					WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), "e", "f", "g", 1).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			actualFn: func(ctx context.Context) error {
				bgr := NewBurjoGroupRepository(s.client)

				bg := domain.BurjoGroup{
					Model: gorm.Model{
						ID:        1,
						CreatedAt: time.Now(),
						UpdatedAt: time.Now(),
						DeletedAt: gorm.DeletedAt{},
					},
					Name:        "e",
					Owner:       "f",
					Description: "g",
				}

				err := bgr.Update(ctx, bg)
				if err != nil {
					return err
				}

				return nil
			},
			wantErr: nil,
		},
		{
			name: "delete",
			expectFn: func() {
				s.mockClient.ExpectExec(`UPDATE "burjo_groups" SET "deleted_at"=$1 WHERE "burjo_groups"."id" = $2 AND "burjo_groups"."deleted_at" IS NULL`).
					WithArgs(sqlmock.AnyArg(), 1).WillReturnResult(sqlmock.NewResult(1, 1))
			},
			actualFn: func(ctx context.Context) error {
				bgr := NewBurjoGroupRepository(s.client)

				bg := domain.BurjoGroup{
					Model: gorm.Model{
						ID:        1,
						CreatedAt: time.Now(),
						UpdatedAt: time.Now(),
						DeletedAt: gorm.DeletedAt{},
					},
					Name:        "a",
					Owner:       "b",
					Description: "c",
				}

				err := bgr.Delete(ctx, bg)
				if err != nil {
					return err
				}

				return nil
			},
			wantErr: nil,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			tc.expectFn()
			err := tc.actualFn(ctx)
			if tc.wantErr == nil {
				assert.NoError(t, err)
			}
		})
	}

	s.Teardown(t)
}
