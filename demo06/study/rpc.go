package study

import (
	"log"
	"net"
	"net/rpc"
)

type HelloService struct{}

// Hello方法必须满足Go语言的RPC规则：方法只能有两个可序列化的参数，其中第二个参数是指针类型，并且返回一个error类型，同时必须是公开的方法
func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "Hello, " + request
	return nil
}

func TestRPC() {
	rpc.RegisterName("HelloService", new(HelloService))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("listen error:", err)
	}
	conn, err := listener.Accept()
	if err != nil {
		log.Fatal("accept error:", err)
	}
	rpc.ServeConn(conn)
}
