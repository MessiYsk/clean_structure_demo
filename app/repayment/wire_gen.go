// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

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

// Injectors from wire.go:

// Module 初始化还款模块
func Module() *ModuleRepayment {
	config := infra.NewConfig()
	db := infra.NewSqlLiteGORM(config)
	creditCardRepository := &persistence.CreditCardRepository{
		DB: db,
	}
	repaymentOrderRepository := &persistence.RepaymentOrderRepository{
		DB: db,
	}
	gormTransactionManager := &persistence.GormTransactionManager{
		DB: db,
	}
	paymentService := &rpc.PaymentService{}
	payoutService := &rpc.PayoutService{}
	repaymentUseCase := &usecase.RepaymentUseCase{
		CreditCardRepository:     creditCardRepository,
		RepaymentOrderRepository: repaymentOrderRepository,
		TransactionManager:       gormTransactionManager,
		PaymentService:           paymentService,
		PayoutService:            payoutService,
	}
	repaymentHandler := &handler.RepaymentHandler{
		RepaymentUseCase: repaymentUseCase,
	}
	usecaseRepaymentUseCase := usecase.RepaymentUseCase{
		CreditCardRepository:     creditCardRepository,
		RepaymentOrderRepository: repaymentOrderRepository,
		TransactionManager:       gormTransactionManager,
		PaymentService:           paymentService,
		PayoutService:            payoutService,
	}
	payoutConsumer := &mq.PayoutConsumer{
		RepaymentUseCase: usecaseRepaymentUseCase,
	}
	paymentConsumer := &mq.PaymentConsumer{
		RepaymentUseCase: usecaseRepaymentUseCase,
	}
	moduleRepayment := &ModuleRepayment{
		RepaymentHandler: repaymentHandler,
		PayoutConsumer:   payoutConsumer,
		PaymentConsumer:  paymentConsumer,
	}
	return moduleRepayment
}

// wire.go:

// infraSet 基础设施依赖集合
//
//go:generate wire
var infraSet = wire.NewSet(infra.NewConfig, infra.NewSqlLiteGORM)

// handlerSet 处理层依赖集合
var handlerSet = wire.NewSet(wire.Struct(new(handler.RepaymentHandler), "*"))

// repoSet 数据库层依赖集合
var repoSet = wire.NewSet(wire.Struct(new(persistence.CreditCardRepository), "db"), wire.Bind(new(repository.ICreditCardRepository), new(*persistence.CreditCardRepository)), wire.Struct(new(persistence.RepaymentOrderRepository), "db"), wire.Bind(new(repository.IRepaymentOrderRepository), new(*persistence.RepaymentOrderRepository)), wire.Struct(new(persistence.GormTransactionManager), "db"), wire.Bind(new(repository.TransactionManager), new(*persistence.GormTransactionManager)))

// rpcSet 外部服务依赖集合
var rpcSet = wire.NewSet(wire.Struct(new(rpc.PaymentService), "*"), wire.Bind(new(usecase.IPaymentService), new(*rpc.PaymentService)), wire.Struct(new(rpc.PayoutService), "*"), wire.Bind(new(usecase.IPayoutService), new(*rpc.PayoutService)))

// adapterSet 适配器层依赖集合
var adapterSet = wire.NewSet(
	handlerSet,
	repoSet,
	rpcSet,
)

// useCaseSet 用例层依赖集合
var useCaseSet = wire.NewSet(wire.Struct(new(usecase.RepaymentUseCase), "*"))

// AppSet 应用层依赖集合
var AppSet = wire.NewSet(
	adapterSet,
	useCaseSet,
)

var consumerSet = wire.NewSet(wire.Struct(new(mq.PaymentConsumer), "*"), wire.Struct(new(mq.PayoutConsumer), "*"))

// ModuleRepayment 还款模块
type ModuleRepayment struct {
	*handler.RepaymentHandler // 还款处理层
	*mq.PayoutConsumer
	*mq.PaymentConsumer
}
