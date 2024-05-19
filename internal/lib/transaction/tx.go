package transaction

import (
	"context"

	"gorm.io/gorm"
)

type txKey struct{}

type Manager interface {
	Execute(ctx context.Context, fn func(context.Context) error) error
}

func NewTransactionTx(ctx context.Context, db *gorm.DB) *gorm.DB {
	return db.Begin()
}

func CheckTransactionFromTx(ctx context.Context) *gorm.DB {
	v, ok := ctx.Value(txKey{}).(*gorm.DB)
	if !ok {
		return nil
	}

	return v
}

func TransactionFromCtx(ctx context.Context, val *gorm.DB) *gorm.DB {
	tx, ok := ctx.Value(txKey{}).(*gorm.DB)
	if !ok {
		return val.WithContext(ctx)
	}

	return tx.WithContext(ctx)
}

func TransactionToCtx(ctx context.Context, db *gorm.DB) context.Context {
	return context.WithValue(ctx, txKey{}, db)
}
