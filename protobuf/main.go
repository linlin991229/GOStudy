package main

import (
	"fmt"

	"google.golang.org/protobuf/proto"
	"lin.com/protobuf_service/proto/userService"
)

func main() {
	info := userService.UserInfo{
		Username: "lin",
		Age:      18,
		Hobby:    []string{"reading", "coding"},
	}

	fmt.Println("info:", info)
	// 默认第一个enum值
	fmt.Println("info.PhoneType:", info.GetType())

	// proto序列化
	data, err := proto.Marshal(&info)
	if err != nil {
		panic(err)
	}
	fmt.Println("data:", data)
	// proto反序列化
	info2 := &userService.UserInfo{}
	err = proto.Unmarshal(data, info2)
	if err != nil {
		panic(err)
	}
	fmt.Println("info2:", info2)
}
