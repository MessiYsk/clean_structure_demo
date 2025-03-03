// Code generated by Kitex v0.12.2. DO NOT EDIT.

package repaymentservice

import (
	"context"
	"errors"
	repayment "github.com/MessiYsk/clean_structure_demo/kitex_gen/repayment"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"ManualRepay": kitex.NewMethodInfo(
		manualRepayHandler,
		newRepaymentServiceManualRepayArgs,
		newRepaymentServiceManualRepayResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	repaymentServiceServiceInfo                = NewServiceInfo()
	repaymentServiceServiceInfoForClient       = NewServiceInfoForClient()
	repaymentServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return repaymentServiceServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return repaymentServiceServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return repaymentServiceServiceInfoForClient
}

// NewServiceInfo creates a new ServiceInfo containing all methods
func NewServiceInfo() *kitex.ServiceInfo {
	return newServiceInfo(false, true, true)
}

// NewServiceInfo creates a new ServiceInfo containing non-streaming methods
func NewServiceInfoForClient() *kitex.ServiceInfo {
	return newServiceInfo(false, false, true)
}
func NewServiceInfoForStreamClient() *kitex.ServiceInfo {
	return newServiceInfo(true, true, false)
}

func newServiceInfo(hasStreaming bool, keepStreamingMethods bool, keepNonStreamingMethods bool) *kitex.ServiceInfo {
	serviceName := "RepaymentService"
	handlerType := (*repayment.RepaymentService)(nil)
	methods := map[string]kitex.MethodInfo{}
	for name, m := range serviceMethods {
		if m.IsStreaming() && !keepStreamingMethods {
			continue
		}
		if !m.IsStreaming() && !keepNonStreamingMethods {
			continue
		}
		methods[name] = m
	}
	extra := map[string]interface{}{
		"PackageName": "repayment",
	}
	if hasStreaming {
		extra["streaming"] = hasStreaming
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.12.2",
		Extra:           extra,
	}
	return svcInfo
}

func manualRepayHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*repayment.RepaymentServiceManualRepayArgs)
	realResult := result.(*repayment.RepaymentServiceManualRepayResult)
	success, err := handler.(repayment.RepaymentService).ManualRepay(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newRepaymentServiceManualRepayArgs() interface{} {
	return repayment.NewRepaymentServiceManualRepayArgs()
}

func newRepaymentServiceManualRepayResult() interface{} {
	return repayment.NewRepaymentServiceManualRepayResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) ManualRepay(ctx context.Context, req *repayment.ManualRepayRequest) (r *repayment.ManualRepayResponse, err error) {
	var _args repayment.RepaymentServiceManualRepayArgs
	_args.Req = req
	var _result repayment.RepaymentServiceManualRepayResult
	if err = p.c.Call(ctx, "ManualRepay", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
