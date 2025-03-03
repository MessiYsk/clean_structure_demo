namespace go repayment

// 信用卡还款请求
struct ManualRepayRequest {
    1: string CreditCardID // 信用卡 ID
    2: double Amount       // 还款金额
    3: double Fee          // 手续费
}

// 信用卡还款响应
struct ManualRepayResponse {
    1: string OrderID     // 订单 ID
    2: string CashierURL  // 收银台链接
    3: string Error       // 错误信息（可选）
}

// 还款服务定义
service RepaymentService {
    ManualRepayResponse ManualRepay(1: ManualRepayRequest req)
}