package main

import (
	"log"
	"net"

	"github.com/cloudwego/kitex/server"
)

// main 程序入口
func main() {

	svr, err := server.InitializeApp()
	if err != nil {
		log.Fatal(err)
	}

	// 启动 Kitex RPC 服务
	addr, _ := net.ResolveTCPAddr("tcp", ":8888")
	rpcServer := repaymentservice.NewServer(&service.RepaymentServiceImpl{repaymentUC: svr.RepaymentUC}, server.WithServiceAddr(addr))

	// 启动消费者

	// 启动 RPC 服务
	err = rpcServer.Run()
	if err != nil {
		log.Fatalf("启动 Kitex 服务失败: %v", err)
	}
}
