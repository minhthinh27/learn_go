package goroutine

import (
	"fmt"
	"sync"
)

func Example4() {
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()

		for i := 100; i >= 0; i-- {
			fmt.Println("100 -> 0", i)
		}
	}()

	go func() {
		defer wg.Done()

		for i := 0; i <= 100; i++ {
			fmt.Println("1 -> 100", i)
		}
	}()

	wg.Wait()

	fmt.Println("Ending")
}
