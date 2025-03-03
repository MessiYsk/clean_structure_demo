package usecase

import (
	"context"

	"github.com/MessiYsk/clean_structure_demo/app/repayment/domain/model"
	"github.com/MessiYsk/clean_structure_demo/app/repayment/domain/repository"
)

// RepaymentUseCase 定义还款用例层
type RepaymentUseCase struct {
	cardRepo   repository.ICreditCardRepository     // 信用卡仓储
	orderRepo  repository.IRepaymentOrderRepository // 还款订单仓储
	txManager  repository.TransactionManager        // 事务管理器
	paymentSvc IPaymentService                      // 支付服务
	payoutSvc  IPayoutService                       // 出款服务
}

// ManualRepayResponse 手动还款响应
type ManualRepayResponse struct {
	OrderID    string // 订单 ID
	CashierURL string // 收银台链接
}

// ManualRepay 处理手动还款请求
func (uc *RepaymentUseCase) ManualRepay(ctx context.Context, creditCardID string, amount float64, fee float64) (*ManualRepayResponse, error) {
	var order *model.RepaymentOrder
	var cashierURL string
	err := uc.txManager.WithTransaction(ctx, func(txCtx context.Context) error {
		// 获取信用卡信息
		card, err := uc.cardRepo.FindByID(creditCardID)
		if err != nil {
			return err
		}

		// 创建初始化状态的还款订单
		order, err = model.NewRepaymentOrder(generateID(), creditCardID, amount+fee)
		if err != nil {
			return err
		}
		if err := uc.orderRepo.Save(order); err != nil {
			return err
		}

		// 调用支付服务获取收银台链接，并更新订单状态为支付中
		cashierURL, err = uc.paymentSvc.CreatePayment(order.ID, amount+fee, card.BankAccountID)
		if err != nil {
			return err
		}
		debitItem, _ := model.NewRepaymentOrderItem(generateID(), order.ID, model.ItemTypeDebit, amount+fee, card.BankAccountID)
		if err := order.StartPayment(*debitItem); err != nil {
			return err
		}
		return uc.orderRepo.Save(order)
	})
	if err != nil {
		return nil, err
	}

	return &ManualRepayResponse{
		OrderID:    order.ID,
		CashierURL: cashierURL,
	}, nil
}

func generateID() string {
	return ""
}
