package dbmodel

import "time"

// RepaymentOrderItem 表示还款订单子单的数据库映射实体
type RepaymentOrderItem struct {
	ID            string    `gorm:"primaryKey"` // 主键：子单 ID
	OrderID       string    `gorm:"index"`      // 所属订单 ID，索引加速查询
	Type          string    // 子单类型
	Amount        float64   // 金额
	AccountID     string    // 关联账户 ID
	Status        string    // 子单状态
	TransactionID string    // 外部交易 ID
	CreatedAt     time.Time // 创建时间
}
