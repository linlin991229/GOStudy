package study

import (
	"fmt"
	"time"
)

func TestChannel5() {

	ch := make(chan int, 3)

	go func() {
		ch <- 1
		ch <- 2
		ch <- 3

		ch <- 4
		fmt.Println("&&&&&")
		ch <- 5
	}()

	go func() {
		for {
			select {
			case n := <-ch:
				time.Sleep(1 * time.Second)
				fmt.Println("n=", n)
			}
		}
	}()

	time.Sleep(6 * time.Second)
}
