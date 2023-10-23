package study

import "fmt"

func TestMap() {
	m := make(map[string]int)
	fmt.Println("map:", m)

	colors := map[string]string{
		"Red":    "#ff0000",
		"Orange": "#ff8000",
		"Yellow": "#ffff00",
	}
	// key可以是任意类型，但必须可以使用==运算符比较
	// 切片、函数以及包含切片的结构类型这些类型由于具有引用语义，不能作为映射的键，使用这些类型会造成编译错误

	value, exists := colors["Blue"]
	// 判断key是否存在
	if !exists {
		fmt.Println("Key does not exists")
	} else {
		fmt.Println("colors[\"Blue\"]:", value)
	}

	if value = colors["Red"]; value != "" {
		fmt.Println("colors[\"Red\"]:", value)
	} else {
		fmt.Println("Key does not exists")
	}

	for k, v := range colors {
		fmt.Println("k:", k, "v:", v)
	}
	delete(colors, "Red")
	fmt.Println("colors:", colors)

	for k := range colors {
		fmt.Println("k:", k)
	}
}
