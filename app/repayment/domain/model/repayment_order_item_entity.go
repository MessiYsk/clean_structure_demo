package model

import (
	"errors"
	"time"
)

// ItemType 表示子单类型
type ItemType string

const (
	ItemTypeDebit  ItemType = "Debit"  // 出款
	ItemTypeCredit ItemType = "Credit" // 入款
	ItemTypeFee    ItemType = "Fee"    // 手续费
)

// ItemStatus 表示子单状态
type ItemStatus string

const (
	ItemStatusProcessing ItemStatus = "Processing" // 处理中
	ItemStatusSuccess    ItemStatus = "Success"    // 成功
	ItemStatusFailed     ItemStatus = "Failed"     // 失败
)

// RepaymentOrderItem 表示还款订单子单（业务实体）
type RepaymentOrderItem struct {
	ID            string     // 子单 ID
	OrderID       string     // 所属订单 ID
	Type          ItemType   // 子单类型
	Amount        float64    // 金额
	AccountID     string     // 关联账户 ID
	Status        ItemStatus // 子单状态
	TransactionID string     // 外部交易 ID
	CreatedAt     time.Time  // 创建时间
}

// NewRepaymentOrderItem 创建处理中的子单
func NewRepaymentOrderItem(id, orderID string, itemType ItemType, amount float64, accountID string) (*RepaymentOrderItem, error) {
	if orderID == "" || accountID == "" || amount <= 0 {
		return nil, errors.New("订单 ID、账户 ID 和金额必须有效")
	}
	return &RepaymentOrderItem{
		ID:        id,
		OrderID:   orderID,
		Type:      itemType,
		Amount:    amount,
		AccountID: accountID,
		Status:    ItemStatusProcessing,
		CreatedAt: time.Now(),
	}, nil
}

// Complete 标记子单成功
func (i *RepaymentOrderItem) Complete(transactionID string) error {
	if i.Status != ItemStatusProcessing {
		return errors.New("子单必须处于处理中状态才能标记成功")
	}
	i.Status = ItemStatusSuccess
	i.TransactionID = transactionID
	return nil
}

// Fail 标记子单失败
func (i *RepaymentOrderItem) Fail() error {
	if i.Status != ItemStatusProcessing {
		return errors.New("子单必须处于处理中状态才能标记失败")
	}
	i.Status = ItemStatusFailed
	return nil
}
