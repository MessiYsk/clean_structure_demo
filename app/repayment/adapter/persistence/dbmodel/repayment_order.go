package dbmodel

import "time"

// RepaymentOrder 表示还款订单的数据库映射实体
type RepaymentOrder struct {
	ID           string    `gorm:"primaryKey"` // 主键：订单 ID
	CreditCardID string    `gorm:"index"`      // 信用卡 ID，索引加速查询
	TotalAmount  float64   // 总金额
	Status       string    // 订单状态
	CreatedAt    time.Time // 创建时间
	UpdatedAt    time.Time // 更新时间
}
