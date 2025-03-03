package persistence

import (
	"context"

	"gorm.io/gorm"
)

// GormTransactionManager GORM 事务管理器
type GormTransactionManager struct {
	db *gorm.DB // 数据库连接
}

// WithTransaction 使用事务包装函数执行逻辑
func (m *GormTransactionManager) WithTransaction(ctx context.Context, fn func(ctx context.Context) error) error {
	tx := m.db.Begin()
	if err := tx.Error; err != nil {
		return err
	}
	txCtx := context.WithValue(ctx, "tx", tx)
	if err := fn(txCtx); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
