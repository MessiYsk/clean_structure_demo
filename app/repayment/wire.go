package repayment

import (
	"clean_structure_demo/app/repayment/adapter/handler"
	"clean_structure_demo/app/repayment/adapter/persistence"
	"clean_structure_demo/app/repayment/adapter/rpc"
	"clean_structure_demo/app/repayment/domain/repository"
	"clean_structure_demo/app/repayment/usecase"
	"clean_structure_demo/infra"

	"github.com/google/wire"
)

//go:generate wire
// infraSet 基础设施依赖集合
var infraSet = wire.NewSet(
	infra.NewConfig, // 配置读取
	infra.NewGORM,   // GORM 数据库连接
)

// handlerSet 处理层依赖集合（对应 Kitex 服务）
var handlerSet = wire.NewSet(
	wire.Struct(new(handler.RepaymentHandler), "*"), // Kitex 还款服务实现
)

// repoSet 数据库层依赖集合
var repoSet = wire.NewSet(
	wire.Struct(new(persistence.CreditCardRepository), "*"),                                  // 信用卡仓储实现
	wire.Bind(new(repository.ICreditCardRepository), new(*persistence.CreditCardRepository)), // 绑定接口到实现

	wire.Struct(new(persistence.RepaymentOrderRepository), "*"),                                      // 还款订单仓储实现
	wire.Bind(new(repository.IRepaymentOrderRepository), new(*persistence.RepaymentOrderRepository)), // 绑定接口到实现

	wire.Struct(new(persistence.GormTransactionManager), "*"),                               // 事务管理器实现
	wire.Bind(new(repository.TransactionManager), new(*persistence.GormTransactionManager)), // 绑定接口到实现
)

// rpcSet 外部服务依赖集合（包括 MQ 和支付/出款服务）
var rpcSet = wire.NewSet(

	wire.Struct(new(rpc.PaymentService), "*"),                         // Mock 支付服务
	wire.Bind(new(usecase.IPaymentService), new(*rpc.PaymentService)), // 绑定支付服务接口到 Mock 实现

	wire.Struct(new(rpc.PayoutService), "*"),                        // Mock 出款服务
	wire.Bind(new(usecase.IPayoutService), new(*rpc.PayoutService)), // 绑定出款服务接口到 Mock 实现
)

// adapterSet 适配器层依赖集合
var adapterSet = wire.NewSet(
	handlerSet, // 处理层
	repoSet,    // 数据库层
	rpcSet,     // 外部服务层
)

// useCaseSet 用例层依赖集合
var useCaseSet = wire.NewSet(
	wire.Struct(new(usecase.RepaymentUseCase), "*"), // 还款用例构造
)

// AppSet 应用层依赖集合
var AppSet = wire.NewSet(
	adapterSet,
	useCaseSet,
)

// ModuleRepayment 还款模块
type ModuleRepayment struct {
	*handler.RepaymentHandler // Kitex 还款服务处理
}

// Module 初始化还款模块
func Module() *ModuleRepayment {
	panic(wire.Build(
		infraSet,
		AppSet,
		wire.Struct(new(ModuleRepayment), "*"),
	))
}
