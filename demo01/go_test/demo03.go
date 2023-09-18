package go_test

import (
	"fmt"
	"math"
	"strings"

	"golang.org/x/tour/wc"
)

type Vertex3 struct {
	X, Y float64
}

var m map[string]Vertex3
var m2 = map[string]Vertex3{
	"Bell Labs": {
		40.68433, -74.39967,
	},
	"Google": {
		37.42202, -122.08408,
	},
}

func mapDemo() {
	m = make(map[string]Vertex3)
	m["Bell Labs"] = Vertex3{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
	fmt.Println(m2)
	delete(m2, "Bell Labs")
	fmt.Println(m2)
	// 若 key 在 m 中，ok 为 true ；否则，ok 为 false。
	// 若 key 不在映射中，那么 elem 是该映射元素类型的零值。
	elem, ok := m2["Bell Labs"]
	fmt.Println("The value:", elem, "Present?", ok)
}
func worldCount(s string) map[string]int {
	words := strings.Fields(s)
	m3 := make(map[string]int)
	for _, v := range words {
		m3[v] = m3[v] + 1
	}
	return m3
}
func practice2() {
	wc.Test(worldCount)
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}
func funV() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))
	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func demo() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(pos(i), neg(-2*i))
	}
}

func fibonacci() func() int {
	x, y := -1, 1
	return func() int {
		sum := x + y
		x, y = y, sum
		return sum
	}
}
func fibonacciDemo() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

// 方法：一类带特殊的 接收者 参数的函数。
func (v Vertex3) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// 方法只是个带接收者参数的函数。
func Abs(v Vertex3) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// 为非结构体类型声明方法，只能为同一个包内定义的类型定义方法
type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}
// 指针接收者声明方法
func (v *Vertex3) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}
func method() {
	v := Vertex3{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(Abs(v))
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
	v.Scale(10)
	fmt.Println(v)
}
