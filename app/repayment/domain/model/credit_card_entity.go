package model

import (
	"errors"
	"time"
)

// CreditCard 表示信用卡账户（业务实体）
type CreditCard struct {
	ID            string    // 信用卡 ID
	UserID        string    // 用户 ID
	CardNumber    string    // 信用卡号
	BankAccountID string    // 默认还款银行账户 ID
	AutoRepay     bool      // 是否启用自动还款
	CreditLimit   float64   // 信用额度
	Balance       float64   // 当前余额（负数表示欠款）
	DueDate       time.Time // 账单到期日
	CreatedAt     time.Time // 创建时间
}

// NewCreditCard 创建新的信用卡账户
func NewCreditCard(id, userID, cardNumber, bankAccountID string, creditLimit float64, dueDate time.Time, autoRepay bool) (*CreditCard, error) {
	if userID == "" || cardNumber == "" || bankAccountID == "" {
		return nil, errors.New("用户 ID、信用卡号和银行账户 ID 不能为空")
	}
	if creditLimit <= 0 {
		return nil, errors.New("信用额度必须为正数")
	}
	if dueDate.Before(time.Now()) {
		return nil, errors.New("到期日必须在未来")
	}
	return &CreditCard{
		ID:            id,
		UserID:        userID,
		CardNumber:    cardNumber,
		BankAccountID: bankAccountID,
		AutoRepay:     autoRepay,
		CreditLimit:   creditLimit,
		Balance:       0,
		DueDate:       dueDate,
		CreatedAt:     time.Now(),
	}, nil
}

// UpdateBalance 更新信用卡余额
func (c *CreditCard) UpdateBalance(amount float64) error {
	newBalance := c.Balance + amount
	if newBalance > 0 {
		return errors.New("信用卡余额不能为正")
	}
	if newBalance < -c.CreditLimit {
		return errors.New("超过信用额度")
	}
	c.Balance = newBalance
	return nil
}
