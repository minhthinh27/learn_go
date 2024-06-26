package goroutine

import (
	"fmt"
)

func Example9() {
	myChan := make(chan int, 1)
	go func() {
		fmt.Printf("%d\n", <-myChan)
	}()
	myChan <- 1
	close(myChan)
}
