package unbuffered

import (
	"fmt"
	"sync"
)

// handle sync data

var x int64 = 0
var increment chan int

func Example1() {
	increment = make(chan int)
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			increment <- 1 // Send 1 to the channel to signal increment
		}()
	}

	go func() {
		for i := range increment {
			x += int64(i) // Receive from the channel and add to x
		}
	}()

	wg.Wait()
	close(increment) // Close the channel after all increments are sent

	fmt.Println("Giá trị của x là: ", x)
}
