package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"lin.com/grpc_server/proto/greeter"
)

type Hello struct{}

func (this Hello) SayHello(c context.Context, req *greeter.HelloReq) (*greeter.HelloRes, error) {
	log.Printf("Hello, %s!", req.Name)
	return &greeter.HelloRes{
		Message: fmt.Sprintf("Hello, %s!", req.Name),
	}, nil
}

func main() {
	// 2.初始grpc对象
	grpcServer := grpc.NewServer()

	// 3.注册服务
	greeter.RegisterGreeterServer(grpcServer, &Hello{})

	// 4.监听端口
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panicln(err)
	}

	// 5.启动服务
	grpcServer.Serve(listener)
}
