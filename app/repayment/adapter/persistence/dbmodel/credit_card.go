package dbmodel

import "time"

// CreditCard 表示信用卡的数据库映射实体
type CreditCard struct {
	ID            string    `gorm:"primaryKey"` // 主键：信用卡 ID
	UserID        string    `gorm:"index"`      // 用户 ID，索引加速查询
	CardNumber    string    `gorm:"unique"`     // 信用卡号，唯一
	BankAccountID string    // 默认还款银行账户 ID
	AutoRepay     bool      // 是否启用自动还款
	CreditLimit   float64   // 信用额度
	Balance       float64   // 当前余额（负数表示欠款）
	DueDate       time.Time // 账单到期日
	CreatedAt     time.Time // 创建时间
}
