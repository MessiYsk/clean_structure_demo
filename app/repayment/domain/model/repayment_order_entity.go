package model

import (
	"errors"
	"time"
)

// OrderStatus 表示还款订单的状态
type OrderStatus string

const (
	StatusInitialized       OrderStatus = "Initialized"       // 初始化
	StatusPaying            OrderStatus = "Paying"            // 支付中
	StatusSuccess           OrderStatus = "Success"           // 支付成功
	StatusPayingOut         OrderStatus = "PayingOut"         // 出款中
	StatusPayoutSuccess     OrderStatus = "PayoutSuccess"     // 出款成功
	StatusFeeTransferred    OrderStatus = "FeeTransferred"    // 手续费转账成功
	StatusFeeTransferFailed OrderStatus = "FeeTransferFailed" // 手续费转账失败
)

// RepaymentOrder 表示还款订单（业务实体）
type RepaymentOrder struct {
	ID           string               // 订单 ID
	CreditCardID string               // 信用卡 ID
	TotalAmount  float64              // 总金额
	Status       OrderStatus          // 订单状态
	Items        []RepaymentOrderItem // 子单列表
	CreatedAt    time.Time            // 创建时间
	UpdatedAt    time.Time            // 更新时间
}

// NewRepaymentOrder 创建初始化状态的还款订单
func NewRepaymentOrder(id, creditCardID string, totalAmount float64) (*RepaymentOrder, error) {
	if creditCardID == "" || totalAmount <= 0 {
		return nil, errors.New("信用卡 ID 和总金额必须有效")
	}
	return &RepaymentOrder{
		ID:           id,
		CreditCardID: creditCardID,
		TotalAmount:  totalAmount,
		Status:       StatusInitialized,
		Items:        []RepaymentOrderItem{},
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}, nil
}

// StartPayment 开始支付，添加支付子单
func (o *RepaymentOrder) StartPayment(debitItem RepaymentOrderItem) error {
	if o.Status != StatusInitialized {
		return errors.New("订单必须处于初始化状态才能开始支付")
	}
	o.Status = StatusPaying
	o.Items = append(o.Items, debitItem)
	o.UpdatedAt = time.Now()
	return nil
}

// ConfirmPayment 确认支付成功
func (o *RepaymentOrder) ConfirmPayment() error {
	if o.Status != StatusPaying {
		return errors.New("订单必须处于支付中状态才能确认支付")
	}
	o.Status = StatusSuccess
	o.UpdatedAt = time.Now()
	return nil
}

// StartPayout 开始出款，添加出款和手续费子单
func (o *RepaymentOrder) StartPayout(creditItem, feeItem RepaymentOrderItem) error {
	if o.Status != StatusSuccess {
		return errors.New("订单必须处于支付成功状态才能开始出款")
	}
	o.Status = StatusPayingOut
	o.Items = append(o.Items, creditItem, feeItem)
	o.UpdatedAt = time.Now()
	return o.validateTotalAmount()
}

// ConfirmPayout 确认出款成功
func (o *RepaymentOrder) ConfirmPayout() error {
	if o.Status != StatusPayingOut {
		return errors.New("订单必须处于出款中状态才能确认出款")
	}
	o.Status = StatusPayoutSuccess
	o.UpdatedAt = time.Now()
	return nil
}

// ConfirmFeeTransfer 确认手续费转账
func (o *RepaymentOrder) ConfirmFeeTransfer(success bool) error {
	if o.Status != StatusPayoutSuccess {
		return errors.New("订单必须处于出款成功状态才能确认手续费转账")
	}
	if success {
		o.Status = StatusFeeTransferred
	} else {
		o.Status = StatusFeeTransferFailed
	}
	o.UpdatedAt = time.Now()
	return nil
}

// validateTotalAmount 校验子单金额总和是否等于订单总金额
func (o *RepaymentOrder) validateTotalAmount() error {
	var sum float64
	for _, item := range o.Items {
		sum += item.Amount
	}
	if sum != o.TotalAmount {
		return errors.New("子单金额总和与订单总金额不匹配")
	}
	return nil
}
