package study

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func TestGoRoutine2() {
	// 分配一个逻辑处理器给调度器
	runtime.GOMAXPROCS(1)

	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 3; i++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 3; i++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	fmt.Println("Waiting To Finish")
	wg.Wait()
	fmt.Println("Done")
}

var wg sync.WaitGroup

func TestGoRoutine3() {

	// runtime.GOMAXPROCS(1)
	runtime.GOMAXPROCS(runtime.NumCPU())

	wg.Add(2)

	// go printPrime("A")
	// go printPrime("B")

	fmt.Println("Waiting To Finish")
	wg.Wait()
	fmt.Println("Done")

}

func printPrime(prefix string) {
	defer wg.Done()

next:
	for i := 0; i < 5000; i++ {
		for j := 2; j < i; j++ {
			if i%j == 0 {
				continue next
			}
		}
		fmt.Printf("%s: %d\n", prefix, i)
	}
	fmt.Println("Completed", prefix)
}

var (
	counter int64
	// shutdown是通知正在执行的goroutine停止工作的标志
	// shutdown int64
	shutdown2 atomic.Int64
)

func TestGoRoutine4() {
	wg.Add(2)

	go incCounter(1)
	go incCounter(2)

	wg.Wait()
	// 存在竞态条件
	fmt.Println("final counter:", counter)
}
func incCounter(id int) {
	defer wg.Done()

	for i := 0; i < 2; i++ {
		value := counter
		fmt.Println("切换前：： id:", id, "counter:", counter)
		// Gosched会让出控制权,切换到其他协程
		runtime.Gosched()
		fmt.Println("id:", id, "counter:", counter)
		value++
		counter = value
		fmt.Println("自增完成：：id:", id, "counter:", counter)
	}
}

func TestGoRoutine5() {
	wg.Add(2)

	go incCounter2(1)
	go incCounter2(2)

	wg.Wait()
	fmt.Println("final counter:", counter)
}

func incCounter2(id int) {
	defer wg.Done()
	for i := 0; i < 2; i++ {
		atomic.AddInt64(&counter, 1)
		runtime.Gosched()
	}
}

func TestGoRoutine6() {
	wg.Add(2)
	go doWork("A")
	go doWork("B")

	time.Sleep(time.Second)
	fmt.Println("Shutdown Now")
	// 安全写
	// atomic.StoreInt64(&shutdown, 1)
	shutdown2.Store(1)
	wg.Wait()
}

func doWork(name string) {
	defer wg.Done()
	for {
		fmt.Printf("Doing %s Work\n", name)
		time.Sleep(250 * time.Millisecond)
		// 安全读
		// if atomic.LoadInt64(&shutdown) == 1 {
		if shutdown2.Load() == 1 {
			fmt.Printf("Shutdown %s Work\n", name)
			break
		}
	}
}

var mutex sync.Mutex

func TestGoRoutine7() {
	wg.Add(2)

	go incCounter3(1)
	go incCounter3(2)
	wg.Wait()
	fmt.Println("final counter:", counter)
}

func incCounter3(id int) {
	defer wg.Done()

	for i := 0; i < 2; i++ {
		// 临界区，同一时刻只允许一个goroutine进入临界区
		mutex.Lock()
		{
			value := counter
			runtime.Gosched()
			value++
			counter = value
		}
		// 释放锁，允许其他goroutine进入
		mutex.Unlock()
	}
}

func TestChannel2() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	court := make(chan int)
	wg.Add(2)

	go player("小明", court)
	go player("小红", court)

	court <- 1

	wg.Wait()
}

func player(name string, court chan int) {
	defer wg.Done()
	for {
		ball, ok := <-court
		if !ok {
			// 如果通道被关闭，我们就赢了
			fmt.Printf("player %s Won\n", name)
			return
		}

		if n := rand.Intn(100); n%13 == 0 {
			fmt.Println("player", name, "被淘汰")
			close(court)
			return
		}
		fmt.Println("player", name, "ball", ball)
		ball++
		court <- ball
	}
}

func TestChannel3() {
	baton := make(chan int)

	wg.Add(1)

	go run(baton)
	// 开始比赛
	baton <- 1
	// 等待比赛结束
	wg.Wait()
}

func run(baton chan int) {
	runner := <-baton

	fmt.Printf("runner: %d running with Baton\n", runner)
	var newRunner int
	if runner != 4 {
		newRunner = runner + 1
		fmt.Printf("runner: %d to the line\n", newRunner)
		go run(baton)
	}
	// 跑
	time.Sleep(100 * time.Millisecond)
	if runner == 4 {
		fmt.Printf("runner: %d finished\n", runner)
		wg.Done()
		return
	}

	fmt.Printf("runner: %d exchange with runner: %d\n", runner, newRunner)
	baton <- newRunner
}

// 无缓冲的通道保证进行发送和接收的goroutine会在同一时间进行数据交换；有缓冲的通道没有这种保证
const (
	numberGoroutine = 4  // 协程数
	taskLoad        = 10 // 任务数
)

func init() {
	// 设置随机数种子，初始化
	rand.New(rand.NewSource(time.Now().Unix()))
}

// TestChannel4
func TestChannel4() {
	tasks := make(chan string, taskLoad)

	wg.Add(numberGoroutine)

	for i := 0; i < numberGoroutine; i++ {
		go worker4(tasks)
	}
	for i := 0; i < taskLoad; i++ {
		tasks <- fmt.Sprintf("Task:%d", i)
	}
	// 所以工作完成后，关闭通道
	close(tasks)
	wg.Wait()
}

func worker4(tasks chan string) {
	defer wg.Done()
	for {
		task, ok := <-tasks
		if !ok {
			fmt.Printf("worker exit\n")
			return
		}
		fmt.Printf("worker: started %s\n", task)
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

		fmt.Printf("worker: done %s\n", task)
	}
}
