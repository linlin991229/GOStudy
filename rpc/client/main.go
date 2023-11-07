package main

import (
	"log"
	"net/rpc"

	"lin.com/rpc_client/rpcInterface"
)

func main() {
	// 1、和rpc微服务建立连接
	conn, err := rpc.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Panicln(err)
	}
	// 3、关闭连接
	defer conn.Close()
	// 2、调用rpc服务
	var result string // 返回值
	conn.Call(rpcInterface.ServerName+".SayHello", "我是客户端client", &result)
	log.Println(result)

	// 4、调用rpc服务 传入结构体参数
	var result2 rpcInterface.Result // 返回值
	conn.Call(rpcInterface.ServerName+".Goods", rpcInterface.GoodsReq{Id: 1, Title: "Mate 60", Price: 8888, Content: "遥遥领先！！！"}, &result2)
	// log.Printf("success: %t, message: %s", result2.Success, result2.Message)
	log.Printf("%+v", result2)

	// 封装rpc服务
	client, err := rpcInterface.DialService("tcp", "127.0.0.1:8080")

	if err != nil {
		log.Panicln(err)
	}
	client.SayHello("我是客户端client封装后的rpc", &result)
	client.Goods(rpcInterface.GoodsReq{Id: 1, Title: "Mate 60", Price: 8888, Content: "遥遥领先！！！"}, &result2)
	log.Println(result)
}
