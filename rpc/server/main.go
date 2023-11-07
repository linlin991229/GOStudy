package main

import (
	"log"
	"net"
	"net/rpc"

	"lin.com/rpc_server/rpc_interface.go"
)

// 1、定义远程调用的服务
type Server struct{}

// 2、定义远程调用的方法
// 方法必须有两个可序列化参数，第二个是指针
// req表示客服端传入的数据，res表示服务端返回的数据
// req和res不能是channel、func，均不能序列化
func (server Server) SayHello(req string, res *string) error {
	*res = "你好" + req
	return nil
}
func (server Server) Goods(req rpc_interface.GoodsReq, res *rpc_interface.Result) error {
	log.Printf("%+v", req)
	*res = rpc_interface.Result{
		Success: true,
		Message: "商品添加成功",
	}
	return nil
}
func registerServer(server rpc_interface.ServiceInterface) error {
	return rpc.RegisterName(rpc_interface.ServerName, server)
}
func main() {
	// 3、注册RPC服务
	if err := registerServer(new(Server)); err != nil {
		log.Panicln(err)
	}
	// 4、监听
	Listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panicln(err)
	}
	log.Println("服务启动成功")
	defer Listener.Close()
	// 5、建立连接
	for {
		conn, err := Listener.Accept()
		if err != nil {
			log.Panicln(err)
		}
		log.Println("客户端连接成功")
		// 6、绑定服务
		go rpc.ServeConn(conn)
	}
}
