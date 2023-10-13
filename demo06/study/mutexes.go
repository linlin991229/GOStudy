package study

import (
	"fmt"
	"sync"
)

type Container struct {
	mu       sync.Mutex
	counters map[string]int
}

func (c *Container) inc(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[key]++
}

func TestMutexes() {
	c := Container{
		counters: map[string]int{"a": 0, "b": 0},
	}
	var wg sync.WaitGroup

	doIncrement := func(name string, n int) {
		for i := 0; i < n; i++ {
			c.inc(name)
		}
		wg.Done()
	}
	wg.Add(3)

	go doIncrement("a", 10000)
	go doIncrement("b", 10000)
	go doIncrement("a", 10000)
	wg.Wait()
	fmt.Println(c.counters)
}
func Test() {
	c := Container{
		counters: make(map[string]int),
	}

	c.counters["a"] = 2
	fmt.Println(c.counters)
	fmt.Println(c.counters["b"])

	m := map[string]int{"a": 0, "b": 0}
	fmt.Println(m)
	m2 := make(map[string]int)
	fmt.Println("m2:", m2)
}
