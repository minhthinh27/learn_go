package goroutine

import (
	"fmt"
	"runtime"
	"sync"
)

func init() {
	runtime.GOMAXPROCS(1)
}

func Example5() {
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 3; i++ {
			for j := 'a'; j < 'z'; j++ {
				fmt.Printf("%c", j)
			}
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 3; i++ {
			for j := 'A'; j < 'Z'; j++ {
				fmt.Printf("%c", j)
			}
		}
	}()

	wg.Wait()

	fmt.Println("Completely")
}
