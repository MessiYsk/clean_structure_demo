package usecase

// IPaymentService 支付服务接口
type IPaymentService interface {
	CreatePayment(orderID string, amount float64, bankAccountID string) (string, error) // 创建支付请求，返回收银台链接
}

// IPayoutService 出款服务接口
type IPayoutService interface {
	CreatePayout(orderID string, creditAmount, feeAmount float64, creditAccountID, feeAccountID string) error // 创建出款请求
}
