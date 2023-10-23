package practice

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func TestIsPrime() {
	for i := 2; i < 100; i++ {
		j := 2
		for ; j < i; j++ {
			if i%j == 0 {
				break
			}
		}
		if j == i {
			fmt.Println("素数:", i)
		}
	}
}
