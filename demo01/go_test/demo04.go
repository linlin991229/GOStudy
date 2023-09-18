package go_test

import (
	"fmt"
	"math"
	"time"
)

// 接口：由一组方法签名定义的集合。
type Abser interface {
	Abs() float64
}

func interfaceDemo() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex3{3, 4}

	a = f // a MyFloat 实现了 Abser
	fmt.Println(a.Abs())
	a = &v // a *Vertex 实现了 Abser

	fmt.Println(a.Abs())

	// 下面一行，v 是一个 Vertex（而不是 *Vertex）
	// 所以没有实现 Abser。
	a = v
	fmt.Println(a.Abs())

	var i I

	i = &T{"Hello"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()

}

type I interface {
	M()
}

type T struct {
	S string
}

// 此方法表示类型 T 实现了接口 I，但我们无需显式声明此事。
func (t T) M() {
	fmt.Println(t.S)
}
func (t *T) M2() {
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}
func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

// 空接口可保存任何类型的值。（因为每个类型都至少实现了零个方法。）

// 空接口被用来处理未知类型的值。例如，fmt.Print 可接受类型为 interface{} 的任意数量的参数。
func emptyInterface() {
	var i interface{}
	describe2(i)

	i = 42
	describe2(i)

	i = "hello"
	describe2(i)
}
func describe2(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
func do(i interface{}) {

	// 类型选择
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

type IPAddr [4]byte

func (ip IPAddr) IPFormat() string {
	return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])
}

func demoStringer() {
	hosts := map[string]IPAddr{
		"loopback": {127, 0, 0, 1},
		"gogleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v:%v\n", name, ip.IPFormat())
	}
}

// 错误
type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("在 %v,%s", e.When, e.What)
}
func run() error {
	return &MyError{
		time.Now(),
		"运行error",
	}
}
func runDemo() {
	// error 为 nil 时表示成功；非 nil 的 error 表示失败。
	if err := run(); err != nil {
		fmt.Println(err)
	}
}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("数据%v不合法", float64(e))
}
func sqrt(x float64) (float64, error) {
	if x < 0 {
		return -1.0, ErrNegativeSqrt(x)
	}
	return math.Sqrt(x), nil
}
func errDemo() {
	fmt.Println(sqrt(2))
	fmt.Println(sqrt(-2))
}
