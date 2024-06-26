package goroutine

import (
	"fmt"
	"time"
)

func Example7() {
	myChan := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			myChan <- i
			time.Sleep(time.Second)
		}
	}()

	for i := 0; i < 10; i++ {
		fmt.Println(<-myChan)
	}

}
