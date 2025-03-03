package infra

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// NewGORM 创建 GORM 数据库连接
func NewGORM(cfg Config) *gorm.DB {
	db, err := gorm.Open(mysql.Open(cfg.MySQL.DSN), &gorm.Config{})
	if err != nil {
		log.Fatalf("连接 MySQL 失败: %v", err)
	}
	return db
}
