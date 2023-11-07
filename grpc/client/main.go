package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"lin.com/grpc_client/proto/greeter"
)

func main() {
	// 1建立连接
	// credentials.NewClientTLSFromFile():从输入的证书文件中为客户端生成TLS凭证。
	// grpc.WithTransportCredentials():设置客户端使用的安全协议,如（TLS/SSL），返回一个客户端安全DialOption的连接。
	grpc_client, err := grpc.Dial("127.0.0.1:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Panicln(err)
	}

	// 2注册客户端
	client := greeter.NewGreeterClient(grpc_client)

	// 3调用服务

	if res, err := client.SayHello(context.Background(), &greeter.HelloReq{Name: "lin"}); err != nil {
		log.Panicln(err)
	} else {
		log.Println(res.Message)

	}
}
