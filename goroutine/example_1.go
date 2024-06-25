package goroutine

import (
	"fmt"
	"runtime"
	"time"
)

func Example1() {
	go func() {
		for i := 1; i <= 50; i++ {
			fmt.Println("I am Goroutine 1")
			runtime.Gosched()
		}
	}()

	go func() {
		for i := 1; i <= 50; i++ {
			fmt.Println("I am Goroutine 2")
			runtime.Gosched()
		}
	}()

	time.Sleep(time.Second)
}
