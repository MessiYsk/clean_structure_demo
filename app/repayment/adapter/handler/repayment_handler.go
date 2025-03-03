package handler

import (
	"clean_structure_demo/kitex_gen/repayment"
	"context"

	"github.com/MessiYsk/clean_structure_demo/app/repayment/usecase"
)

// RepaymentHandler Kitex 还款服务实现
type RepaymentHandler struct {
	repaymentUseCase *usecase.RepaymentUseCase // 还款用例层依赖
}

// ManualRepay 处理手动还款请求
func (s *RepaymentHandler) ManualRepay(ctx context.Context, req *repayment.ManualRepayRequest) (*repayment.ManualRepayResponse, error) {
	resp, err := s.repaymentUseCase.ManualRepay(ctx, req.CreditCardID, req.Amount, req.Fee)
	if err != nil {
		return &repayment.ManualRepayResponse{Error: err.Error()}, nil
	}
	return &repayment.ManualRepayResponse{
		OrderID:    resp.OrderID,
		CashierURL: resp.CashierURL,
	}, nil
}
