package main

import (
	"fmt"
	"math"
	"math/rand"
	"mymodule/go_test"
	"runtime"
	"time"
	"rsc.io/quote"
)

// 函数外的每个语句都必须以关键字开始（var, func 等等）
var c, python, java bool
var i, j int = 1, 2

// 常量
const Pi = 3.14

func main() {
	fmt.Println(c, python, java, i, j)
	var cPlus, rust, golang = true, false, "go"
	fmt.Println(cPlus, rust, golang)
	k := 1
	fmt.Printf("Type: %T,Value: %v \n", k, k)

	fmt.Println(Pi)

	go_test.OutShow()
	fmt.Println(go_test.Add(1, 99))
	fmt.Println(time.Now())
	fmt.Println(rand.Intn(10))
	fmt.Println(math.Pi)

	a, b := go_test.Swap("a", "b")
	fmt.Println(a, b)
	x, y := go_test.Split(10)
	fmt.Println(x, y)

	go_test.Forloop()
	go_test.Whileloop()
	fmt.Printf("go_test.Sqrt(-1): %v\n", go_test.Sqrt(-1))

	fmt.Printf("go_test.Pow(3, 2, 100): %v\n", go_test.Pow(3, 2, 100))
	fmt.Printf("go_test.Pow(5, 2, 20): %v\n", go_test.Pow(5, 2, 20))
	fmt.Println(go_test.NewtonSqrt(1000), math.Sqrt(5))

	go_test.MySwitch()

	// defer 语句会将函数推迟到外层函数返回之后执行。
	// 推迟调用的函数其参数会立即求值，但直到外层函数返回前该函数都不会被调用。
	defer fmt.Println("world")
	fmt.Println("hello")
	go_test.DeferKey()



	go_test.Study01()
	go_test.Study02()

	go_test.Study03()
	fmt.Println(runtime.Version())
	fmt.Println(quote.Go())
}
