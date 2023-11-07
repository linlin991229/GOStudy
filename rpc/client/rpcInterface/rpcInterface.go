package rpcInterface

import "net/rpc"

const ServerName = "Server"

type GoodsReq struct {
	Id      int
	Title   string
	Price   float32
	Content string
}

type Result struct {
	Success bool
	Message string
}
type ServerClient struct {
	*rpc.Client
}

func DialService(network, addr string) (*ServerClient, error) {
	client, err := rpc.Dial(network, addr)
	if err != nil {
		return nil, err
	}
	return &ServerClient{client}, nil
}

func (c *ServerClient) SayHello(name string, reply *string) error {
	return c.Client.Call(ServerName+".SayHello", name, reply)
}

func (c *ServerClient) Goods(req GoodsReq, reply *Result) error {
	return c.Client.Call(ServerName+".Goods", req, reply)
}
