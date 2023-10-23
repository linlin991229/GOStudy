package practice

import "fmt"

func TestChannelClose() {
	ch := make(chan int, 4)
	ch <- 1
	fmt.Println("通道长度:", len(ch))
	// 已关闭的通道任何可以接收数据，但是不能再发送数据
	close(ch)
	fmt.Println("通道长度:", len(ch))
	for i := 0; i < 4; i++ {
		v, ok := <-ch
		fmt.Println("值:", v, "是否关闭:", ok)
	}
}
