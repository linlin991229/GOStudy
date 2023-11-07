package rpc_interface

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
type ServiceInterface interface {
	SayHello(string, *string) error
	Goods(GoodsReq, *Result) error
}
