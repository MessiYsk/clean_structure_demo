package rpc

// PaymentService 定义支付服务接口
type PaymentService struct {
}

// CreatePayment 创建支付请求，返回收银台链接
func (p *PaymentService) CreatePayment(orderID string, amount float64, bankAccountID string) (string, error) {
	// TODO 调用下游支付服务接口
	return "", nil
}
