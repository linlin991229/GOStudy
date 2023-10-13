package study

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// jobs <-chan int: 是一个只读的通道，用于接收工作任务的任务队列。
// results chan<- int: 是一个只写的通道，用于向结果队列发送工作任务的结果。
// results chan int: 是一个普通通道，既可以发送数据也可以接收数据。
// 发送者-接收者模式
func worker2(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func TestWorker() {
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)
	// 启动3个goroutine
	for w := 1; w <= 3; w++ {
		go worker2(w, jobs, results)
	}
	// 发送5个任务
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	// 关闭jobs
	close(jobs)
	// 接收所有处理结果，确保所有的任务都被处理
	for a := 1; a <= numJobs; a++ {
		fmt.Println("received", <-results)
	}
}
func worker3(id int) {
	fmt.Println("worker", id, "started")
	time.Sleep(time.Second)
	fmt.Println("worker", id, "finished")
}
func TestWaitGroup() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		// 循环中引用迭代变量，使用延迟函数、go语句会有并发问题
		// 创建一个新变量解决这个问题
		// 在 Go 语言中，块级作用域内的变量会隐藏外部同名变量，这意味着内部的 i 变量会遮蔽外部的 i 变量。
		// 外部的 i 变量是 for 循环的迭代变量，它在每次循环迭代中都会改变其值。

		// 内部的 i 变量是通过 i := i 重新声明的局部变量，它只在当前循环迭代的块级作用域内有效。

		// 内部的 i 变量遮蔽了外部的 i 变量，意味着在内部的块级作用域中，你可以访问并修改内部的 i 变量，而不会影响外部的 i 变量。

		i := i

		go func() {
			defer wg.Done()
			worker3(i)
		}()

	}
	wg.Wait()
}
func TestRateLimit() {
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)
	limiter := time.Tick(200 * time.Millisecond)

	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}
	burstyLimiter := make(chan time.Time, 3)

	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}

func TestAtomic() {
	var ops atomic.Uint64

	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 1000; j++ {
				ops.Add(1)
			}
			wg.Done()

		}()
	}
	wg.Wait()

	fmt.Println("ops:", ops.Load())
}
