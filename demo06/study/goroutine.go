package study

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}
func TestGoRoutine() {
	f("direct")

	go f("goroutine")
	// 匿名调用
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("done")
}
func TestChannel() {
	message := make(chan string)

	go func() {
		message <- "ping"
	}()
	msg := <-message

	fmt.Println(msg)
}
func TestBufferedChannel() {
	message := make(chan string, 2)
	message <- "buffered"
	message <- "channel"

	fmt.Println("received", <-message)
	fmt.Println("received", <-message)
}

func worker(done chan bool) {
	fmt.Println("working...")
	time.Sleep(time.Second)
	fmt.Println("done")
	done <- true
}

func TestWorkerPool() {
	// 创建一个bool型的channel
	done := make(chan bool, 1)
	go worker(done)
	// 阻塞等待worker结束
	<-done
}

// 该ping函数仅接受用于发送值的通道。尝试在此通道上接收将是一个编译时错误
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// 该pong函数接受一个通道用于接收 ( pings)，第二个通道用于发送 ( pongs)。
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func TestChannelDirections() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}

func TestSelect() {
	c1 := make(chan string)
	c2 := make(chan string)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "one"
	}() // 匿名函数de调用，立即执行
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		// 使用select语句可以同时监听（等待）多个通道
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}

func TestSelect2() {
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()

	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}
}

func TestSelect3() {
	message := make(chan string)
	signal := make(chan bool)

	select {
	// 如果没有可用值，直接执行default语句
	case msg := <-message:
		fmt.Println("received", msg)
	default:
		fmt.Println("no message received")
	}
	msg := "hi"
	select {
	// 此通道没有缓冲区，也没有接受者，直接执行default语句
	case message <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}
	// 多路非阻塞select
	select {
	case msg := <-message:
		fmt.Println("received message", msg)
	case sig := <-signal:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}

func TestCloseChannel() {
	jobs := make(chan int, 3)
	done := make(chan bool)
	go func() {
		for {
			// ok用于判断是否接收成功，ok为true表示接收成功，可以避免通道关闭接收到0值
			// num, ok := <-c
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()
	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")
	<-done
}
func TestRangeChannel() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	// 可以关闭非空通道，但仍然可以接收剩余的值。
	close(queue)
	for elem := range queue {
		fmt.Println(elem)
	}
}
