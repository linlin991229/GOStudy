package study

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func TestConcurrent() {
	ch := make(chan string)

	go func() {
		ch <- searchByBing("golang")
	}()

	go func() {
		ch <- searchByGoogle("golang")
	}()
	go func() {
		ch <- searchByBaidu("golang")
	}()
	// 开启冗余goroutine只采用最先完成的goroutine
	select {
	case res := <-ch:
		fmt.Println("completed:", res)
	}
}

func searchByBing(word string) string {
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	return "bing completed"
}
func searchByGoogle(word string) string {
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	return "google completed"
}
func searchByBaidu(word string) string {
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	return "baidu completed"
}

var file *os.File

func TestConcurrentPrime() {
	defer file.Close()
	ch := generateNature()
	for i := 0; i < 100; i++ {
		prime := <-ch
		fmt.Println("第:", i+1, " prime:", prime)
		ch = primeFilter(ch, prime, i+1)

	}
}
func generateNature() chan int {
	ch := make(chan int)

	go func() {
		for i := 2; ; i++ {
			ch <- i
		}
	}()
	return ch
}

// in <-chan int: 是一个只读的通道，用于接收工作任务的任务队列。
func primeFilter(in <-chan int, prime int, id int) chan int {
	file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		panic(err)
	}
	out := make(chan int)
	go func() {
		for {
			i := <-in
			fmt.Fprintf(file, "协程id:%d,接收到的数:%d,素数:%d\n", id, i, prime)
			if i%prime != 0 {
				out <- i
			}
		}
	}()
	return out
}
