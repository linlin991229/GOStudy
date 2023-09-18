package go_test

import (
	"fmt"
	"math"
	"math/rand"
	"runtime"
	"strings"
	"time"

	"golang.org/x/tour/pic"
)

func Add(a, b int) int {
	return a + b
}

// 多值返回
func Swap(x, y string) (string, string) {
	return y, x
}

// 命名返回 没有参数的 return 语句返回已命名的返回值。也就是 直接 返回。
func Split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func Forloop() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

// 无条件就是无限循环
func Whileloop() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

func Sqrt(x float64) string {
	if x < 0 {
		return Sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func Pow(x, n, lim float64) float64 {
	// 该语句声明的变量作用域仅在 if 之内。同样可以在任何对应的 else 块中使用
	if pow := math.Pow(x, n); pow < lim {
		return pow
	} else {
		fmt.Println(pow, lim)
	}
	return lim
}

func NewtonSqrt(x float64) float64 {
	z := x / 2
	// z :=float64(1)
	for i := 0; i < 10; i++ {
		beforeV := z
		z -= (z*z - x) / (2 * z)
		fmt.Println("第", i, "次猜测", z)
		if math.Abs(z-beforeV) <= 0.000001 {
			fmt.Println("获得近似值", z)
			return z
		}
	}
	return z
}

func MySwitch() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}

	t := time.Now()
	// 没有条件的 switch 同 switch true 一样
	//这种形式能将一长串 if-then-else 写得更加清晰。
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

func DeferKey() {
	fmt.Println("counting")
	defer fmt.Println("done")
	// 存储在栈中，后进先出的顺序
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
}

func point() {
	i, j := 42, 2701
	p := &i
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)

	p = &j
	*p = *p / 37
	fmt.Println(j)
}

type Vertex struct {
	X int
	Y int
}
type Vertex2 struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}  // 创建一个 Vertex 类型的结构体
	v2 = Vertex{X: 1}  // Y:0 被隐式地赋予
	v3 = Vertex{}      // X:0 Y:0
	p  = &Vertex{1, 2} // 创建一个 *Vertex 类型的结构体（指针）
)

func struct01() {
	vertex := Vertex{1, 2}
	fmt.Println(vertex)
	fmt.Println(vertex.X)
	p := &vertex
	p.X = 1e9
	fmt.Println(vertex)

}
func arr() {
	var a [2]string
	a[0] = "小"
	a[1] = "霖霖"
	fmt.Println(a[0], a)
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	// 切片 为数组元素提供动态大小，更灵活
	var s []int = primes[1:4] // 指下标，[1,4)
	fmt.Println(s)
	// 切片并不存储任何数据，它只是描述了底层数组中的一段。

	// 更改切片的元素会修改其底层数组中对应的元素。

	// 与它共享底层数组的切片都会观测到这些修改。

	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a1 := names[0:2]
	b := names[1:3]
	fmt.Println(a1, b)

	b[0] = "XXX"
	fmt.Println(a1, b)
	fmt.Println(names)

	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s1 := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s1)
	// 切片下界的默认值为 0，上界则是该切片的长度。
	s2 := []int{2, 3, 5, 7, 11, 13}

	s2 = s2[1:4]
	fmt.Println(s2)

	s2 = s2[:2]
	fmt.Println(s2)

	s2 = s2[1:]
	fmt.Println(s2)

	s3 := []int{2, 3, 5, 7, 11, 13}
	printSlice(s3)

	// 截取切片使其长度为 0
	s3 = s3[:0]
	printSlice(s3)

	// 拓展其长度，不会改变容量cap
	s3 = s3[:4]
	printSlice(s3)

	// 舍弃前两个值，会改变容量cap，会变小
	s3 = s3[2:]
	printSlice(s3)
	// 切片的零值是 nil。
	s3 = s3[1:]
	printSlice(s3)
	// 切片的零值是 nil。
}
func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
func slice() {
	// 创建切片
	a := make([]int, 5)
	printSlice(a)
	// 指定容量
	b := make([]int, 0, 5)
	printSlice(b)
}
func boardFunc() {
	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"
	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

func appendSlice() {
	var s []int
	printSlice(s)

	s = append(s, 0)
	printSlice(s)

	s = append(s, 1)
	printSlice(s)

	s = append(s, 2, 3, 4)
	printSlice(s)
}

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func rangeFor() {
	// 第一个值为当前元素的下标，第二个值为该下标所对应元素的一份副本。
	for i, v := range pow {
		fmt.Printf("2**%d=%d\n", i, v)
	}

	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}
func practice() {
	pic.Show(Pic)
}
func Pic(dx, dy int) [][]uint8 {
	pic := make([][]uint8, dy)
	for i := 0; i < len(pic); i++ {
		pic[i] = make([]uint8, dx)
		for j := 0; j < dx; j++ {
			switch rand.Intn(5) {
			case 0:
				pic[i][j] = uint8(dx+dy) / 2
			case 1:
				pic[i][j] = uint8(dx * dy)
			case 2:
				pic[i][j] = uint8(dx ^ dy)
			case 3:
				pic[i][j] = uint8(float64(dx) * math.Log(float64(dy)))
			case 4:
				pic[i][j] = uint8(dx % (dy + 1))
			}
		}
	}
	fmt.Println(pic)
	return pic
}
