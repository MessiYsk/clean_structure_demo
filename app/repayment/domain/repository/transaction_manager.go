package repository

import "context"

// TransactionManager GORM 事务管理器
type TransactionManager interface {
	WithTransaction(ctx context.Context, fn func(ctx context.Context) error) error
}
