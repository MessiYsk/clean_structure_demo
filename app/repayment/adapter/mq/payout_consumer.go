package mq

import "github.com/MessiYsk/clean_structure_demo/app/repayment/usecase"

//PayoutConsumer 出款结果通知(出款)
type PayoutConsumer struct {
	RepaymentUseCase usecase.RepaymentUseCase
}

func (receiver *PayoutConsumer) Handle() {
	// TODO
}
