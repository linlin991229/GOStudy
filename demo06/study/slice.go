package study

import "fmt"

// func TestSlice() {
// 	// 长度和容量相等
// 	slice := make([]string, 5)
// 	// 指定长度和容量
// 	slice2 := make([]string, 5, 10)

// 	// 长度和容量相等为5
// 	slice3 := []string{"a", "b", "c", "d", "e"}

// 	// 使用空字符串初始化第 100 个元素
// 	slice4 := []string{99: ""}
// 	// 如果在[]运算符里指定了一个值，那么创建的就是数组而不是切片。只有不指定值的时候，才会创建切片，
// 	array := [3]string{"a", "b", "c"}

// 	// 创建nil切片
// 	var slice5 []string

// 	// 空切片
// 	slice6 := make([]string, 0)
// 	slice7 := []int{}

// }

func TestSlice2() {
	// 从切片中获取元素
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// 由于地址指针保存的是第一个元素的地址，所以下面切片后失去了第一个元素的地址，容量也减小了
	newSlice := slice[1:4]
	fmt.Println("newSlice:", newSlice)
	// 对底层数组的容量是k的切片slice[i:j]来说
	// 长度：j-i
	// 容量: k-i
	// 此时slice与newSlice共享内存区域，如果一个切片修改了该底层数组的共享部分，另一个切片也能感知到
	newSlice[0] = 100
	fmt.Println("slice:", slice)

	// 增加切片的长度
	newSlice = append(newSlice, 100)
	fmt.Println("newSlice:", newSlice)
	// 此时slice相应的底层数组也会发生变化
	fmt.Println("slice:", slice)

	// 达到切片的容量但小于1000时，成倍增长，超过1000时，变为1.25倍（增加25%）
	// 此时newSlice2是一个新的底层数组，不会影响slice
	newSlice2 := append(slice, 200)
	newSlice2[0] = 100
	fmt.Println("newSlice2:", newSlice2)
	fmt.Println("slice:", slice)

	// 对于slice[i:j:K]
	// 长度：j-i
	// 容量: k-i

	// 如果在创建切片时设置切片的容量和长度一样，就可以强制让新切片的第一个append操作创建新的底层数组
	// 这样可以与原有的底层数组分离，不会影响原有的底层数组
	newSlice3 := slice[1:4:4]
	fmt.Println("newSlice3:", newSlice3)
	newSlice3 = append(newSlice3, 1000)
	fmt.Println("newSlice3:", newSlice3)
	fmt.Println("slice:", slice)

	s1 := []int{1, 2}
	s2 := []int{3, 4}
	// 最佳可变长度参数
	s3 := append(s1, s2...)
	fmt.Println(s3)

	// 迭代
	// range会创建副本
	for i, v := range slice {
		fmt.Printf("Index: %d Value: %d\n", i, v)
	}
	for i := 2; i < len(slice); i++ {
		fmt.Printf("Index: %d Value: %d\n", i, slice[i])
	}

	// 多维切片
	s4 := [][]int{{1}, {3, 4}}
	fmt.Println("s4:", s4)
	s4[0] = append(s4[0], 2)
	fmt.Println("s4:", s4)
}
