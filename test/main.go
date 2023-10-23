package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(100)
	// 循环变量问题
	for i := 0; i < 100; i++ {
		go func(i int) {
			defer wg.Done()
			fmt.Println("i=", i)
		}(i)
	}
	wg.Wait()
	fmt.Println("over")
}
