package mq

import "github.com/MessiYsk/clean_structure_demo/app/repayment/usecase"

//PaymentConsumer 支付结果通知（入款)
type PaymentConsumer struct {
	RepaymentUseCase usecase.RepaymentUseCase
}

func (receiver *PaymentConsumer) Handle() {
	// TODO
}
