//go:build wireinject
// +build wireinject

package repayment

import (
	"github.com/MessiYsk/clean_structure_demo/app/repayment/adapter/handler"
	"github.com/MessiYsk/clean_structure_demo/app/repayment/adapter/mq"
	"github.com/MessiYsk/clean_structure_demo/app/repayment/adapter/persistence"
	"github.com/MessiYsk/clean_structure_demo/app/repayment/adapter/rpc"
	"github.com/MessiYsk/clean_structure_demo/app/repayment/domain/repository"
	"github.com/MessiYsk/clean_structure_demo/app/repayment/usecase"
	"github.com/MessiYsk/clean_structure_demo/infra"
	"github.com/google/wire"
)

//go:generate wire
// infraSet 基础设施依赖集合
var infraSet = wire.NewSet(
	infra.NewConfig,      // 配置读取
	infra.NewSqlLiteGORM, // GORM 数据库连接
)

// handlerSet 处理层依赖集合
var handlerSet = wire.NewSet(
	wire.Struct(new(handler.RepaymentHandler), "*"), // Kitex 或其他处理层实现，注入所有字段
)

// repoSet 数据库层依赖集合
var repoSet = wire.NewSet(
	wire.Struct(new(persistence.CreditCardRepository), "db"),                                 // 信用卡仓储实现
	wire.Bind(new(repository.ICreditCardRepository), new(*persistence.CreditCardRepository)), // 绑定接口

	wire.Struct(new(persistence.RepaymentOrderRepository), "db"),                                     // 还款订单仓储实现
	wire.Bind(new(repository.IRepaymentOrderRepository), new(*persistence.RepaymentOrderRepository)), // 绑定接口

	wire.Struct(new(persistence.GormTransactionManager), "db"),                              // 事务管理器实现
	wire.Bind(new(repository.TransactionManager), new(*persistence.GormTransactionManager)), // 绑定接口
)

// rpcSet 外部服务依赖集合
var rpcSet = wire.NewSet(
	wire.Struct(new(rpc.PaymentService), "*"),                         // 支付服务实现（Mock 或真实）
	wire.Bind(new(usecase.IPaymentService), new(*rpc.PaymentService)), // 绑定支付服务接口

	wire.Struct(new(rpc.PayoutService), "*"),                        // 出款服务实现（Mock 或真实）
	wire.Bind(new(usecase.IPayoutService), new(*rpc.PayoutService)), // 绑定出款服务接口
)

// adapterSet 适配器层依赖集合
var adapterSet = wire.NewSet(
	handlerSet, // 处理层
	repoSet,    // 数据库层
	rpcSet,     // 外部服务层
)

// useCaseSet 用例层依赖集合
var useCaseSet = wire.NewSet(
	wire.Struct(new(usecase.RepaymentUseCase), "*"), // 还款用例构造，注入所有字段
)

// AppSet 应用层依赖集合
var AppSet = wire.NewSet(
	adapterSet,
	useCaseSet,
)

var consumerSet = wire.NewSet(
	wire.Struct(new(mq.PaymentConsumer), "*"),
	wire.Struct(new(mq.PayoutConsumer), "*"),
)

// ModuleRepayment 还款模块
type ModuleRepayment struct {
	*handler.RepaymentHandler // 还款处理层
	*mq.PayoutConsumer
	*mq.PaymentConsumer
}

// Module 初始化还款模块
func Module() *ModuleRepayment {
	panic(wire.Build(
		infraSet,
		AppSet,
		consumerSet,
		wire.Struct(new(ModuleRepayment), "*"),
	))
}
