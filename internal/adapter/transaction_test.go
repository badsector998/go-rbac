package adapter

import (
	"context"
	"sync"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/badsector998/go-rbac/internal/domain"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm/schema"
)

func TestTransaction(t *testing.T) {
	ts := new(TestSuite)
	ts.Setup(t)

	// uncomment to include read from db
	bgModel, err := schema.Parse(&domain.BurjoGroup{}, &sync.Map{}, schema.NamingStrategy{})
	assert.NoError(t, err)

	cols := make([]string, 0)
	for _, i := range bgModel.Fields {
		cols = append(cols, i.DBName)
	}

	tt := []struct {
		name     string
		ctx      context.Context
		expectFn func()
		actualFn func(ctx context.Context) error
		wantErr  error
	}{
		{
			name: "success insert tx",
			ctx:  context.Background(),
			expectFn: func() {
				ts.txMockClient.ExpectBegin()
				ts.txMockClient.ExpectQuery(`INSERT INTO "burjo_groups" ("created_at","updated_at","deleted_at","name","owner","description") VALUES ($1,$2,$3,$4,$5,$6) RETURNING "id"`).
					WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), "a", "b", "c").
					WillReturnRows(sqlmock.NewRows([]string{"id"}))
				ts.txMockClient.ExpectCommit()
			},
			actualFn: func(ctx context.Context) error {
				bgRepo := NewBurjoGroupRepository(ts.txClient)
				err := ts.tx.Execute(ctx, func(ctx context.Context) error {
					bg := domain.BurjoGroup{
						Name:        "a",
						Owner:       "b",
						Description: "c",
					}

					err := bgRepo.Create(ctx, bg)
					if err != nil {
						return err
					}

					return nil
				})

				return err
			},
			wantErr: nil,
		},
		{
			name: "success retrieve all burjo groups",
			ctx:  context.Background(),
			expectFn: func() {
				ts.mockClient.ExpectQuery(`SELECT * FROM "burjo_groups" WHERE "burjo_groups"."deleted_at" IS NULL`).
					WillReturnRows(sqlmock.NewRows(cols))
			},
			actualFn: func(ctx context.Context) error {
				bg := NewBurjoGroupRepository(ts.client)
				_, err := bg.ReadAll(ctx)
				return err
			},
			wantErr: nil,
		},
		{
			name: "success to get burjo group with ID 1",
			ctx:  context.Background(),
			expectFn: func() {
				ts.mockClient.ExpectQuery(`SELECT * FROM "burjo_groups" WHERE id = $1 AND "burjo_groups"."deleted_at" IS NULL`).
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows(cols))
			},
			actualFn: func(ctx context.Context) error {
				bg := NewBurjoGroupRepository(ts.client)
				_, err := bg.ReadByID(ctx, 1)
				return err
			},
			wantErr: nil,
		},
		{
			name: "sucess to get then update within transaction",
			ctx:  context.Background(),
			expectFn: func() {
				ts.txMockClient.ExpectBegin()
				ts.txMockClient.ExpectQuery(`SELECT * FROM "burjo_groups" WHERE id = $1 AND "burjo_groups"."deleted_at" IS NULL`).
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows(cols))
				ts.txMockClient.ExpectExec(
					`UPDATE "burjo_groups" SET "created_at"=$1,"updated_at"=$2,"deleted_at"=$3,"name"=$4,"owner"=$5,"description"=$6 WHERE "burjo_groups"."deleted_at" IS NULL AND "id" = $7`,
				).WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), "e", "f", "g", 1).WillReturnResult(sqlmock.NewResult(1, 1))
				ts.txMockClient.ExpectCommit()
			},
			actualFn: func(ctx context.Context) error {
				bg := NewBurjoGroupRepository(ts.client)
				err := ts.tx.Execute(ctx, func(ctx context.Context) error {
					b, err := bg.ReadByID(ctx, 1)
					if err != nil {
						return err
					}

					b.ID = 1
					b.Name = "e"
					b.Owner = "f"
					b.Description = "g"

					err = bg.Update(ctx, b)
					if err != nil {
						return err
					}

					return nil
				})

				return err
			},
			wantErr: nil,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			tc.expectFn()
			err := tc.actualFn(tc.ctx)
			assert.NoError(t, err)
		})
	}

	ts.Teardown(t)
}
