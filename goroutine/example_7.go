package goroutine

import (
	"fmt"
)

func Example7() {
	myChan := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			myChan <- i
		}
	}()

	for i := 0; i < 10; i++ {
		fmt.Println(<-myChan)
	}

}
