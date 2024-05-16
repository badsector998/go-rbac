package adapter

import (
	"context"
	"testing"

	"github.com/badsector998/go-rbac/internal/domain"
	"github.com/badsector998/go-rbac/internal/lib/transaction"
)

func TestTransaction(t *testing.T) {
	ts := new(TestSuite)
	ts.Setup(t)

	bg := domain.BurjoGroup{
		Name:        "a",
		Owner:       "b",
		Description: "c",
	}

	ts.txMockClient.ExpectBegin()
	ts.txMockClient.ExpectExec("INSERT INTO burjo_groups(name, owner, description) VALUES('a', 'b', 'c')").WillReturnError(nil)
	ts.txMockClient.ExpectExec("UPDATE burjo_groups SET name='e', owner='f', description='g' WHERE ID=1").WillReturnError(nil)
	ts.txMockClient.ExpectCommit()

	ctx := context.Background()

	err := ts.tx.Execute(ctx, func(ctx context.Context) error {
		c := transaction.TransactionFromCtx(ctx, ts.txClient)
		err := c.Save(&bg).Error
		if err != nil {
			return err
		}

		var bgs domain.BurjoGroup
		err = c.Model(&domain.BurjoGroup{}).First(&bgs).Error
		if err != nil {
			return err
		}

		bgs.Name = "e"
		bgs.Owner = "f"
		bgs.Description = "g"

		err = c.Save(&bgs).Error
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		t.Errorf("error: %s", err.Error())
	}

	ts.Teardown(t)
}
