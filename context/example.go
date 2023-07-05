package context

import (
	"context"
	"fmt"
	"time"
)

func Run() {
	taskfb := func(ctx context.Context) <-chan int {
		dts := make(chan int)
		n := 1

		go func() {
			for {
				select {
				case <-ctx.Done():
					fmt.Println("exe cannel")
					return
				case dts <- n:
					fmt.Println("case dts <- n:")
					n++
				}
			}
		}()
		return dts
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for n := range taskfb(ctx) {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}

	time.Sleep(time.Second * 3)
}
