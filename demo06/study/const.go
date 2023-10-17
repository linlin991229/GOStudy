package study

import (
	"fmt"
)

// iota递增关键字
const (
	a = iota + 1
	b
	c
	d
)

func TestIterate() {
	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("c:", c)
	fmt.Println("d:", d)
}
