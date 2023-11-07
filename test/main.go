package main

import (
	"fmt"
	"math/rand"
	"sync"
)

// 匿名字段
type Student struct {
	string
}

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
	temp := make(map[int]string)
	temp[1] = "value"
	fmt.Println(temp)

	fmt.Println(rand.Intn(10))

	x, y := 1, 2
	x += y
	y = x - y
	x -= y
	fmt.Println(x, y)
}
