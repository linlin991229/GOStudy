package study

import "fmt"

func TestArray() {
	// 长度由元素个数决定
	arr := [...]int{1, 2, 3, 4} // []int{1,2,3,4}是切片
	fmt.Println(arr)
	strs := []string{"aaa", "bbb"}
	fmt.Println(strs)
	// 字符串转换为字节切片
	bytes := []byte("a b")
	fmt.Println(bytes)

	var array1 [3]*string
	array2 := [3]*string{new(string), new(string), new(string)}

	*array2[0] = "Red"
	*array2[1] = "Green"
	*array2[2] = "Blue"

	array1 = array2
	fmt.Println("array1:", array1)

	// 多维数组
	// 普通声明
	// var array [4][2]int
	// // 通过字面量声明并初始化
	// array3 := [4][2]int{
	// 	{1, 2},
	// 	{3, 4},
	// 	{5, 6},
	// 	{7, 8},
	// }
	// 声明并初始化外层数组中索引为1个和3的元素
	array4 := [4][2]int{1: {20, 21}, 3: {30, 31}}
	fmt.Println("array4:", array4)
	// 声明并初始化外层数组和内层数组的单个元素
	array5 := [4][2]int{1: {0: 20}, 3: {1: 41}}
	fmt.Println("array5:", array5)

	// 数组作为函数参数传递时使用指针效率更高，但要避免改变指针指向
}
