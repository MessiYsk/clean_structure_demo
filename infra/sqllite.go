package infra

import (
	"github.com/MessiYsk/clean_structure_demo/app/repayment/adapter/persistence/dbmodel"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

// NewSqlLiteGORM 创建 GORM 数据库连接（SQLite）
func NewSqlLiteGORM(cfg Config) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(cfg.SQLite.Path), &gorm.Config{})
	if err != nil {
		log.Fatalf("连接 SQLite 失败: %v", err)
	}
	// 自动迁移数据库表
	err = db.AutoMigrate(
		&dbmodel.CreditCard{},
		&dbmodel.RepaymentOrder{},
		&dbmodel.RepaymentOrderItem{},
	)
	if err != nil {
		log.Fatalf("自动迁移 SQLite 表失败: %v", err)
	}
	return db
}
