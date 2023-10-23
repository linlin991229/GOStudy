package study

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func worker5(ctx context.Context, wg *sync.WaitGroup) error {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("worker5 Done!!!")
			return ctx.Err()
		default:
			fmt.Println("worker5")
		}
	}
}
func TestContext() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go worker5(ctx, &wg)
	}
	time.Sleep(time.Second)
	cancel()

	wg.Wait()

}
