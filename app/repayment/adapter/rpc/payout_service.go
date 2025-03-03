package rpc

type PayoutService struct {
}

// CreatePayout 创建出款请求
func (p *PayoutService) CreatePayout(orderID string, creditAmount, feeAmount float64, creditAccountID, feeAccountID string) error {
	return nil
}
