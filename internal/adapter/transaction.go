package adapter

import (
	"context"
	"fmt"

	"github.com/badsector998/go-rbac/internal/lib/transaction"
	"gorm.io/gorm"
)

type TransactionAdapter struct {
	db *gorm.DB
}

func NewTransactionAdapter(db *gorm.DB) *TransactionAdapter {
	return &TransactionAdapter{db: db}
}

var _ transaction.Manager = &TransactionAdapter{}

func (t *TransactionAdapter) Execute(ctx context.Context, fn func(context.Context) error) error {
	tx := transaction.CheckTransactionFromTx(ctx)

	if tx == nil {
		tx = transaction.NewTransactionTx(ctx, t.db)
	}

	if tx.Error != nil {
		return fmt.Errorf("failed to start transaction: %w", tx.Error)
	}

	ctx = transaction.TransactionToCtx(ctx, tx)

	err := fn(ctx)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to commit. rolling back: %w", err)
	}

	err = tx.Commit().Error
	if err != nil {
		return fmt.Errorf("failed to commit. error on commiting: %w", err)
	}

	return nil
}
