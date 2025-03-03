package repository

import (
	"github.com/MessiYsk/clean_structure_demo/app/repayment/domain/model"
)

// IRepaymentOrderRepository 定义还款订单仓储接口
type IRepaymentOrderRepository interface {
	Save(order *model.RepaymentOrder) error                                  // 保存还款订单
	FindByID(id string) (*model.RepaymentOrder, error)                       // 按 ID 查询还款订单
	FindByCreditCardID(creditCardID string) ([]*model.RepaymentOrder, error) // 按信用卡 ID 查询订单列表
}
