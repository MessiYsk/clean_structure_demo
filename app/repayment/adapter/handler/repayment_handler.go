package handler

import (
	"context"

	"github.com/MessiYsk/clean_structure_demo/app/repayment/usecase"
	"github.com/MessiYsk/clean_structure_demo/kitex_gen/repayment"
)

// RepaymentHandler Kitex 还款服务实现
type RepaymentHandler struct {
	RepaymentUseCase *usecase.RepaymentUseCase // 还款用例层依赖
}

// ManualRepay 处理手动还款请求
func (s *RepaymentHandler) ManualRepay(ctx context.Context, req *repayment.ManualRepayRequest) (*repayment.ManualRepayResponse, error) {
	resp, err := s.RepaymentUseCase.ManualRepay(ctx, req.CreditCardID, req.Amount, req.Fee)
	if err != nil {
		return &repayment.ManualRepayResponse{Error: err.Error()}, nil
	}
	return &repayment.ManualRepayResponse{
		OrderID:    resp.OrderID,
		CashierURL: resp.CashierURL,
	}, nil
}
